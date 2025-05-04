package account

import (
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock/account_mock.go -package=mock -source=./account_interface.go

// Event represents a domain event
type Event interface {
	// GetID returns the unique identifier of the event
	GetID() uuid.UUID

	// GetContextID returns the unique identifier of the context that the event belongs to, i.e. account ID, customer ID, etc.
	GetContextID() uuid.UUID

	// GetOrigin returns the origin of the event, i.e. account, customer, etc.
	GetOrigin() string

	// GetType returns the type of the event, i.e. account.funds.withdrawn, customer.created, etc.
	GetType() string

	// GetTypeVersion returns the version of the event's type, i.e. 0.0.1
	GetTypeVersion() string

	// GetState returns the state of the event, i.e. created, completed, failed, aborted
	GetState() string

	// GetCreatedAt returns the date and time the event was created
	GetCreatedAt() time.Time

	// GetScheduledAt returns the date and time the event was scheduled to be processed (if applicable)
	GetScheduledAt() time.Time

	// GetStartedAt returns the date and time the event was started
	GetStartedAt() time.Time

	// GetCompletedAt returns the date and time the event was completed
	GetCompletedAt() time.Time

	// GetRetry returns the number of times the event has been processed, field set to 1 when the event is scheduled, incremented when the event is retried up to MaxRetry
	GetRetry() int

	// GetMaxRetry returns the maximum number of times the event can be retried
	GetMaxRetry() int

	// GetEventData returns the event data
	GetEventData() []byte
}
