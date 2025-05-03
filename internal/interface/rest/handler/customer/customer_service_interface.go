package customer

import (
	"context"

	customerapplication "github.com/stefanowiczd/ddd-case-01/internal/application/customer"
)

//go:generate mockgen -destination=./mock/customer_service_mock.go -package=mock -source=./customer_service_interface.go

// CustomerQueryService defines the contract for customer query service that handles queries/read operations
type CustomerQueryService interface {
	// GetCustomer gets a customer by ID
	GetCustomer(ctx context.Context, dto customerapplication.GetCustomerDTO) (customerapplication.GetCustomerResponseDTO, error)
}

// CustomerService defines the contract for the customer service that handles commands/mutable operations
type CustomerService interface {
	// CreateCustomer creates a new customer
	CreateCustomer(ctx context.Context, dto customerapplication.CreateCustomerDTO) (customerapplication.CreateCustomerResponseDTO, error)
	// UpdateCustomer updates a customer
	UpdateCustomer(ctx context.Context, dto customerapplication.UpdateCustomerDTO) error
	// BlockCustomer blocks a customer
	BlockCustomer(ctx context.Context, dto customerapplication.BlockCustomerDTO) error
	// UnblockCustomer unblocks a customer
	UnblockCustomer(ctx context.Context, dto customerapplication.UnblockCustomerDTO) error
	// DeleteCustomer deletes a customer
	DeleteCustomer(ctx context.Context, dto customerapplication.DeleteCustomerDTO) error
}
