package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	customerapplication "github.com/stefanowiczd/ddd-case-01/internal/application/customer"
)

// Handler handles HTTP requests for customer operations
type CustomerHandler struct {
	customerService CustomerService
}

func NewCustomerHandler(customerService CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

type CreateCustomerRequest struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type CreateCustomerResponse struct {
	Customer Customer `json:"customer"`
}

type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address"`
}

func (r *CreateCustomerRequest) Validate() error {
	// TODO: validate fields
	return nil
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.CreateCustomer(
		r.Context(),
		customerapplication.CreateCustomerDTO{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Address: customerapplication.Address{
				Street:     req.Address.Street,
				City:       req.Address.City,
				State:      req.Address.State,
				PostalCode: req.Address.PostalCode,
				Country:    req.Address.Country,
			},
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(customer) // TODO decide about handling of this error.
}

type UpdateCustomerRequest struct {
	CustomerID string
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Address    Address `json:"address"`
}

func (r *UpdateCustomerRequest) Validate() error {
	// TODO: validate fields
	return nil
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var req UpdateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.CustomerID = r.PathValue("id")

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.customerService.UpdateCustomer(
		r.Context(),
		customerapplication.UpdateCustomerDTO{
			CustomerID: req.CustomerID,
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			Email:      req.Email,
			Phone:      req.Phone,
			Address: customerapplication.Address{
				Street:     req.Address.Street,
				City:       req.Address.City,
				State:      req.Address.State,
				PostalCode: req.Address.PostalCode,
				Country:    req.Address.Country,
			},
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type BlockCustomerRequest struct {
	CustomerID string
	Reason     string `json:"reason"`
}

func (r *BlockCustomerRequest) Validate() error {
	if _, err := uuid.Parse(r.CustomerID); err != nil {
		return fmt.Errorf("validate: customer id as uuid: %w", err)
	}

	if r.Reason == "" {
		return fmt.Errorf("validate: reason is required")
	}

	return nil
}

func (h *CustomerHandler) BlockCustomer(w http.ResponseWriter, r *http.Request) {
	var req BlockCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.CustomerID = r.PathValue("customerId")

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.customerService.BlockCustomer(
		r.Context(),
		customerapplication.BlockCustomerDTO{
			CustomerID: req.CustomerID,
			Reason:     req.Reason,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CustomerHandler) UnblockCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerId")

	if _, err := uuid.Parse(customerID); err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	err := h.customerService.UnblockCustomer(
		r.Context(),
		customerapplication.UnblockCustomerDTO{
			CustomerID: customerID,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerId")

	if _, err := uuid.Parse(customerID); err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	err := h.customerService.DeleteCustomer(
		r.Context(),
		customerapplication.DeleteCustomerDTO{
			CustomerID: customerID,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
