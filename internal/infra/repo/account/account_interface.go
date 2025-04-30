package account

import (
	"time"

	"github.com/google/uuid"
)

// Event represents a domain event
type AccountEvent interface {
	// GetID returns the unique identifier of the event
	GetID() uuid.UUID

	// GetType returns the type of the event
	GetType() string

	// GetTypeVersion returns the version of the event's type
	GetTypeVersion() string

	// GetAggregateID returns the unique identifier of the aggregate that the event belongs to
	GetAggregateID() uuid.UUID

	// GetCreatedAt returns the date and time the event was created
	GetCreatedAt() time.Time

	// GetEventData returns the event data
	GetEventData() []byte

	// GetScheduledAt returns the earliest date and time the event can be scheduled
	GetScheduledAt() time.Time
}
