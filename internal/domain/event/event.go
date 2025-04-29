package event

import (
	"time"

	"github.com/google/uuid"
)

// Event represents a domain event
type Event interface {
	GetID() string
	GetType() string
	GetAggregateID() string
	GetCreatedAt() time.Time
}

// BaseEvent provides common fields for all domain events
type BaseEvent struct {
	ID          uuid.UUID
	Type        string
	AggregateID uuid.UUID
	CreatedAt   time.Time
}

func (e BaseEvent) GetID() uuid.UUID {
	return e.ID
}

func (e BaseEvent) GetType() string {
	return e.Type
}

func (e BaseEvent) GetAggregateID() uuid.UUID {
	return e.AggregateID
}

func (e BaseEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
