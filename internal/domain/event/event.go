package event

import (
	"time"
)

// Event represents a domain event
type Event interface {
	EventID() string
	EventType() string
	AggregateID() string
	Timestamp() time.Time
}

// BaseEvent provides common fields for all domain events
type BaseEvent struct {
	ID             string
	Type           string
	Aggregate      string
	EventTimestamp time.Time
}

func (e BaseEvent) EventID() string {
	return e.ID
}

func (e BaseEvent) EventType() string {
	return e.Type
}

func (e BaseEvent) AggregateID() string {
	return e.Aggregate
}

func (e BaseEvent) Timestamp() time.Time {
	return e.EventTimestamp
}
