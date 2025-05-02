package customer

import (
	"context"

	"github.com/google/uuid"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

//go:generate mockgen -destination=./mock/customer_service_mock.go -package=mock -source=./customer_service_interface.go

//             Customer

// CustomerQueryRepository defines the interface for customer queries
type CustomerQueryRepository interface {
	// FindByID retrieves a customer by its ID
	FindByID(ctx context.Context, id uuid.UUID) (*customerdomain.Customer, error)
}

// CustomerEventRepository defines the interface for customer event persistence
type CustomerEventRepository interface {
	// CreateEvents persists a customer event
	CreateEvents(ctx context.Context, events []customerdomain.Event) error
}
