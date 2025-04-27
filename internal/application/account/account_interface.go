package account

import (
	"context"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

//go:generate mockgen -destination=./mock/account_service_mock.go -package=mock -source=./account_interface.go

// AccountQueryRepository defines the interface for account queries
type AccountQueryRepository interface {
	// Save persists an account
	Save(ctx context.Context, account *accountdomain.Account) error

	// FindByID retrieves an account by its ID
	FindByID(ctx context.Context, id string) (*accountdomain.Account, error)

	// FindByAccountNumber retrieves an account by its account number
	FindByAccountNumber(ctx context.Context, accountNumber string) (*accountdomain.Account, error)

	// FindByCustomerID retrieves all accounts by a customer ID
	FindByCustomerID(ctx context.Context, customerID string) ([]*accountdomain.Account, error)

	// Update updates an existing account
	Update(ctx context.Context, account *accountdomain.Account) error

	// Delete removes an account
	Delete(ctx context.Context, id string) error
}

// CustomerQueryRepository defines the interface for customer queries
type CustomerQueryRepository interface {
	// FindByID retrieves a customer by its ID
	FindByID(ctx context.Context, id string) (*customerdomain.Customer, error)
}

// AccountEventRepository defines the interface for account event persistence
type AccountEventRepository interface {
	// CreateEvent persists an account event
	CreateEvents(ctx context.Context, events []accountdomain.Event) error
}
