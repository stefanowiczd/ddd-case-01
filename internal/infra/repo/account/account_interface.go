package account

import (
	"time"

	"github.com/google/uuid"
)

// BaseEvent represents a domain event
type BaseEvent interface {
	// GetID returns the unique identifier of the event
	GetID() uuid.UUID

	// GetAccountID returns the unique identifier of the account that the event belongs to
	GetAccountID() uuid.UUID

	// GetAggregateID returns the unique identifier of the aggregate that the event belongs to
	GetAggregateID() uuid.UUID

	// GetType returns the type of the event
	GetType() string

	// GetTypeVersion returns the version of the event's type
	GetTypeVersion() string

	// GetState returns the state of the event
	GetState() string

	// GetCreatedAt returns the date and time the event was created
	GetCreatedAt() time.Time

	GetCompletedAt() time.Time

	// GetScheduledAt returns the earliest date and time the event can be scheduled
	GetScheduledAt() time.Time

	// GetRetry returns the number of times the event has been retried
	GetRetry() int

	// GetMaxRetry returns the maximum number of times the event can be retried
	GetMaxRetry() int

	// GetEventData returns the event data
	GetEventData() []byte
}
