package processor

import (
	"context"

	"github.com/google/uuid"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
	eventdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

//go:generate mockgen -destination=./mock/processor_mock.go -package=mock -source=./processor_interface.go

type BaseEvent interface {
	GetID() uuid.UUID
	GetOrigin() string
	GetType() string
	GetEventData() []byte
}

// OrchestratorRepository defines the interface for event orchestration operations
type OrchestratorRepository interface {
	// Query Operations

	// FindAllEvents returns all events
	FindAllEvents(ctx context.Context) ([]*eventdomain.BaseEvent, error)
	// FindProcessableEvents returns all events that are ready to be processed
	FindProcessableEvents(ctx context.Context, limit int) ([]*eventdomain.BaseEvent, error)
	// FindByOriginAndStatus returns all events that match the origin and status
	FindByOriginAndStatus(ctx context.Context, origin, state string, limit int) ([]*eventdomain.BaseEvent, error)
	// FindByID returns an event by its ID
	FindByID(ctx context.Context, id uuid.UUID) (*eventdomain.BaseEvent, error)

	// Command Operations
	// UpdateEventStart updates the event at start up
	UpdateEventStart(ctx context.Context, id uuid.UUID) error
	// UpdateEventCompletion updates the event at completion
	UpdateEventCompletion(ctx context.Context, id uuid.UUID) error
	// UpdateEventRetry updates the event at retry
	UpdateEventRetry(ctx context.Context, id uuid.UUID, retry int) error
	// UpdateEventState updates the event state
	UpdateEventState(ctx context.Context, id uuid.UUID, state string) error
}

// AccountRepository defines the interface for account operations
type AccountRepository interface {
	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, account *accountdomain.Account) error
	// GetAccount returns an account by its ID
	FindByID(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error)
}

// CustomerRepository defines the interface for customer operations
type CustomerRepository interface {
	// CreateCustomer creates a new customer
	CreateCustomer(ctx context.Context, customer *customerdomain.Customer) error
}
