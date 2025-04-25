package event

import (
	"time"
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
	ID          string
	Type        string
	AggregateID string
	CreatedAt   time.Time
}

func (e BaseEvent) GetID() string {
	return e.ID
}

func (e BaseEvent) GetType() string {
	return e.Type
}

func (e BaseEvent) GetAggregateID() string {
	return e.AggregateID
}

func (e BaseEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
