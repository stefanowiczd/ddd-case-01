package customer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

var (
	// ErrCustomerNotFound is returned when an customer is not found.
	ErrCustomerNotFound = errors.New("customer not found")
	// ErrCustomerAlreadyExists is returned when a customer already exists.
	ErrCustomerAlreadyExists = errors.New("customer already exists")
)

type Customer = customerdomain.Customer
type Address = customerdomain.Address

func ToCustomerDTO(customer *Customer) CustomerResponseDTO {
	return CustomerResponseDTO{
		ID:        customer.ID.String(),
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Address:   customer.Address,
		Status:    customer.Status.String(),
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

// CustomerService handles customer-related use cases
type CustomerService struct {
	customerQueryRepo CustomerQueryRepository

	customerEventRepo CustomerEventRepository
}

func NewCustomerService(customerQueryRepo CustomerQueryRepository, customerEventRepo CustomerEventRepository) *CustomerService {
	return &CustomerService{
		customerQueryRepo: customerQueryRepo,
		customerEventRepo: customerEventRepo,
	}
}

type CreateCustomerDTO struct {
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	DateOfBirth string
	Address     Address
}

type CreateCustomerResponseDTO struct {
	Customer CustomerResponseDTO
}

// CreateCustomer creates a new customer
func (c *CustomerService) CreateCustomer(ctx context.Context, dto CreateCustomerDTO) (CreateCustomerResponseDTO, error) {
	customerID := uuid.New()

	_, err := c.customerQueryRepo.FindByID(ctx, customerID)
	if err != nil && !errors.Is(err, customerdomain.ErrCustomerNotFound) {
		return CreateCustomerResponseDTO{}, fmt.Errorf("finding customer by id: %w", err)
	}

	customer := customerdomain.NewCustomer(customerID, dto.FirstName, dto.LastName, dto.Phone, dto.Email, dto.DateOfBirth, dto.Address)

	err = c.customerEventRepo.CreateEvents(ctx, customer.Events)
	if err != nil {
		return CreateCustomerResponseDTO{}, fmt.Errorf("creating customer events: %w", err)
	}

	return CreateCustomerResponseDTO{Customer: ToCustomerDTO(customer)}, nil
}

type CustomerResponseDTO struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   Address
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetCustomerDTO struct {
	CustomerID string
}

type GetCustomerResponseDTO struct {
	Customer CustomerResponseDTO `json:"customer"`
}

// GetCustomer retrieves a customer by its ID
func (c *CustomerService) GetCustomer(ctx context.Context, dto GetCustomerDTO) (GetCustomerResponseDTO, error) {
	customer, err := c.customerQueryRepo.FindByID(ctx, uuid.MustParse(dto.CustomerID))
	if err != nil {
		if errors.Is(err, customerdomain.ErrCustomerNotFound) {
			return GetCustomerResponseDTO{}, fmt.Errorf("finding customer by id: %w", ErrCustomerNotFound)
		}

		return GetCustomerResponseDTO{}, fmt.Errorf("finding customer by id: %w", err)
	}

	return GetCustomerResponseDTO{
		Customer: ToCustomerDTO(customer),
	}, nil
}

type UpdateCustomerDTO struct {
	CustomerID  string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	DateOfBirth string
	Address     Address
}

func (c *CustomerService) detectChanges(dto UpdateCustomerDTO) customerdomain.CustomerEventType {
	var eventType customerdomain.CustomerEventType

	if dto.FirstName != "" || dto.LastName != "" || dto.Email != "" || dto.Phone != "" {
		eventType = customerdomain.CustomerUpdatedAllEventType
	}

	if dto.FirstName != "" && dto.LastName != "" {
		eventType = customerdomain.CustomerUpdatedNameEventType
	}

	if dto.Email != "" || dto.Phone != "" {
		eventType = customerdomain.CustomerUpdatedContactEventType
	}

	if dto.Address.Street != "" || dto.Address.City != "" || dto.Address.State != "" || dto.Address.PostalCode != "" || dto.Address.Country != "" {
		eventType = customerdomain.CustomerUpdatedAddressEventType
	}

	return eventType
}

func (c *CustomerService) UpdateCustomer(ctx context.Context, dto UpdateCustomerDTO) error {
	customer, err := c.customerQueryRepo.FindByID(ctx, uuid.MustParse(dto.CustomerID))
	if err != nil {
		if errors.Is(err, customerdomain.ErrCustomerNotFound) {
			return fmt.Errorf("finding customer by id: %w", ErrCustomerNotFound)
		}

		return fmt.Errorf("finding customer by id: %w", err)
	}

	updateEventType := c.detectChanges(dto)

	customer.Update(updateEventType, dto.FirstName, dto.LastName, dto.Phone, dto.Email, dto.DateOfBirth, dto.Address)

	err = c.customerEventRepo.CreateEvents(ctx, customer.Events)
	if err != nil {
		return fmt.Errorf("creating customer events: %w", err)
	}

	return nil
}

type BlockCustomerDTO struct {
	CustomerID string
	Reason     string
}

func (c *CustomerService) BlockCustomer(ctx context.Context, dto BlockCustomerDTO) error {
	customer, err := c.customerQueryRepo.FindByID(ctx, uuid.MustParse(dto.CustomerID))
	if err != nil {
		if errors.Is(err, customerdomain.ErrCustomerNotFound) {
			return fmt.Errorf("finding customer by id: %w", ErrCustomerNotFound)
		}

		return fmt.Errorf("finding customer by id: %w", err)
	}

	customer.Block(dto.Reason)

	err = c.customerEventRepo.CreateEvents(ctx, customer.Events)
	if err != nil {
		return fmt.Errorf("creating customer events: %w", err)
	}

	return nil
}

type UnblockCustomerDTO struct {
	CustomerID string
}

func (c *CustomerService) UnblockCustomer(ctx context.Context, dto UnblockCustomerDTO) error {
	customer, err := c.customerQueryRepo.FindByID(ctx, uuid.MustParse(dto.CustomerID))
	if err != nil {
		if errors.Is(err, customerdomain.ErrCustomerNotFound) {
			return fmt.Errorf("finding customer by id: %w", ErrCustomerNotFound)
		}

		return fmt.Errorf("finding customer by id: %w", err)
	}

	customer.Unblock()

	err = c.customerEventRepo.CreateEvents(ctx, customer.Events)
	if err != nil {
		return fmt.Errorf("creating customer events: %w", err)
	}

	return nil
}

type DeleteCustomerDTO struct {
	CustomerID string
}

func (c *CustomerService) DeleteCustomer(ctx context.Context, dto DeleteCustomerDTO) error {
	customer, err := c.customerQueryRepo.FindByID(ctx, uuid.MustParse(dto.CustomerID))
	if err != nil {
		if errors.Is(err, customerdomain.ErrCustomerNotFound) {
			return nil
		}

		return fmt.Errorf("finding customer by id: %w", err)
	}

	customer.Delete()

	err = c.customerEventRepo.CreateEvents(ctx, customer.Events)
	if err != nil {
		return fmt.Errorf("creating customer events: %w", err)
	}

	return nil
}
