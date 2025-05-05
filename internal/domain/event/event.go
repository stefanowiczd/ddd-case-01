package event

import (
	"time"

	"github.com/google/uuid"
)

// BaseEvent provides common fields for all domain events
type BaseEvent struct {
	// ID is the unique identifier for the event
	ID uuid.UUID `json:"id"`
	// ContextID is the ID of the context that the event belongs to, e.g. account ID, customer ID, etc.
	ContextID uuid.UUID `json:"context_id"`
	// Origin is the origin of the event, e.g. account, customer, etc.
	Origin string `json:"origin"`
	// Type is the type of the event, e.g. account.funds.withdrawn, customer.created, etc.
	Type string `json:"type"`
	// TypeVersion is the version of the event type, e.g. 0.0.1
	TypeVersion string `json:"type_version"`
	// State is the state of the event, e.g. created, completed, failed, aborted
	State string `json:"state"`
	// CreatedAt is the time the event was created
	CreatedAt time.Time `json:"created_at"`
	// ScheduledAt is the time the event was scheduled to be processed (if applicable)
	ScheduledAt time.Time `json:"scheduled_at"`
	// StartedAt is the time the event was started
	StartedAt time.Time `json:"started_at"`
	// CompletedAt is the time the event was completed
	CompletedAt time.Time `json:"completed_at"`
	// Retry is the number of times the event has been processed, field set to 1 when the event is scheduled, incremented when the event is retried up to MaxRetry
	Retry int `json:"retry"`
	// MaxRetry is the maximum number of times the event an be retried
	MaxRetry int `json:"max_retry"`
	// Data is the data associated with the event
	Data []byte `json:"data"`
}

func NewBaseEvent(
	id, contextID uuid.UUID,
	origin, typeEvent, typeVersion string, createdAt, scheduledAt time.Time, maxRetry int) BaseEvent {
	return BaseEvent{
		ID:          id,
		ContextID:   contextID,
		Origin:      origin,
		Type:        typeEvent,
		TypeVersion: typeVersion,
		State:       EventStateCreated.String(),
		CreatedAt:   createdAt,
		ScheduledAt: scheduledAt,
		Retry:       0,
		MaxRetry:    maxRetry,
		Data:        nil,
	}
}

func (e *BaseEvent) GetID() uuid.UUID {
	return e.ID
}

func (e *BaseEvent) GetContextID() uuid.UUID {
	return e.ContextID
}

func (e *BaseEvent) GetOrigin() string {
	return e.Origin
}

func (e *BaseEvent) GetType() string {
	return e.Type
}

func (e *BaseEvent) GetTypeVersion() string {
	return e.TypeVersion
}

func (e *BaseEvent) GetState() string {
	return e.State
}

func (e *BaseEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *BaseEvent) GetScheduledAt() time.Time {
	return e.ScheduledAt
}

func (e *BaseEvent) GetStartedAt() time.Time {
	return e.StartedAt
}

func (e *BaseEvent) GetCompletedAt() time.Time {
	return e.CompletedAt
}

func (e *BaseEvent) GetRetry() int {
	return e.Retry
}

func (e *BaseEvent) GetMaxRetry() int {
	return e.MaxRetry
}

func (e *BaseEvent) GetEventData() []byte {
	return e.Data
}

func (e *BaseEvent) Schedule(t time.Time) {
	e.ScheduledAt = t
}
