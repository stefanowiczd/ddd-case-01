package customer

import (
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=./mock/customer_mock.go -package=mock -source=./customer_interface.go

// Event represents a domain event
type Event interface {
	// GetID returns the unique identifier of the event
	GetID() uuid.UUID

	// GetContextID returns the unique identifier of the context that the event belongs to, i.e. account ID, customer ID, etc.
	GetContextID() uuid.UUID

	// GetType returns the type of the event
	GetType() string

	// GetTypeVersion returns the version of the event's type
	GetTypeVersion() string

	// GetState returns the state of the event
	GetState() string

	// GetCreatedAt returns the date and time the event was created
	GetCreatedAt() time.Time

	// GetCompletedAt returns the date and time the event was completed
	GetCompletedAt() time.Time

	// GetScheduledAt is the time the event was scheduled to be processed (if applicable)
	GetScheduledAt() time.Time

	// GetRetry returns the number of times the event has been processed, field set to 1 when the event is scheduled, incremented when the event is retried up to MaxRetry
	GetRetry() int

	// GetMaxRetry returns the maximum number of times the event can be retried
	GetMaxRetry() int

	// GetEventData returns the event data
	GetEventData() []byte
}
