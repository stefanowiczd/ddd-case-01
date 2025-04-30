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
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	TypeVersion string    `json:"type_version"`
	AggregateID uuid.UUID `json:"aggregate_id"`
	CreatedAt   time.Time `json:"created_at"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Data        []byte    `json:"data"`
}

func NewBaseEvent(id uuid.UUID, t, tv string, aggregateID uuid.UUID, scheduledAt time.Time) BaseEvent {
	return BaseEvent{
		ID:          id,
		Type:        t,
		TypeVersion: tv, // TODO add retry logic around this
		AggregateID: aggregateID,
		CreatedAt:   time.Now().UTC(),
		ScheduledAt: scheduledAt, // TODO add retry logic around this
		Data:        nil,
	}
}

func (e BaseEvent) GetID() uuid.UUID {
	return e.ID
}

func (e BaseEvent) GetType() string {
	return e.Type
}

func (e BaseEvent) GetTypeVersion() string {
	return e.TypeVersion
}

func (e BaseEvent) GetAggregateID() uuid.UUID {
	return e.AggregateID
}

func (e BaseEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e BaseEvent) GetEventData() []byte {
	return e.Data
}

func (e BaseEvent) GetScheduledAt() time.Time {
	return e.ScheduledAt
}

func (e *BaseEvent) Schedule(t time.Time) {
	e.ScheduledAt = t
}
