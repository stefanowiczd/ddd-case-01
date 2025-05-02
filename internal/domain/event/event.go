package event

import (
	"time"

	"github.com/google/uuid"
)

// EventState represents the state of an event
type EventState string

// String returns the string representation of the event state
func (e EventState) String() string {
	return string(e)
}

const (
	// EventStateCreated is the state of the event when it is created
	EventStateCreated EventState = "created"
	// EventStateCreatedRepeated is the state of the event when it is created and repeated
	EventStateCreatedRepeated EventState = "created.repeated"
	// EventStateCompleted is the state of the event when it is completed
	EventStateCompleted EventState = "completed"
	// EventStateFailed is the state of the event when it is failed after all retries
	EventStateFailed EventState = "failed"
	// EventStateAborted is the state of the event when it is aborted, i.e. by the user
	EventStateAborted EventState = "aborted"
)

// BaseEvent provides common fields for all domain events
type BaseEvent struct {
	// ID is the unique identifier for the event
	ID uuid.UUID `json:"id"`
	// ContextID is the ID of the context that the event belongs to, e.g. account ID, customer ID, etc.
	ContextID uuid.UUID `json:"context_id"`
	// Type is the type of the event
	Type string `json:"type"`
	// TypeVersion is the version of the event type
	TypeVersion string `json:"type_version"`
	// State is the state of the event
	State string `json:"state"`
	// CreatedAt is the time the event was created
	CreatedAt time.Time `json:"created_at"`
	//StartedAt??
	// CompletedAt is the time the event was completed
	CompletedAt time.Time `json:"completed_at"`
	// ScheduledAt is the time the event was scheduled to be processed (if applicable)
	ScheduledAt time.Time `json:"scheduled_at"`
	// Retry is the number of times the event has been processed, field set to 1 when the event is scheduled, incremented when the event is retried up to MaxRetry
	Retry int `json:"retry"`
	// MaxRetry is the maximum number of times the event an be retried
	MaxRetry int `json:"max_retry"`
	// Data is the data associated with the event
	Data []byte `json:"data"`
}

func NewBaseEvent(id, contextID uuid.UUID, t, tv string, scheduledAt time.Time, maxRetry int) BaseEvent {
	return BaseEvent{
		ID:          id,
		ContextID:   contextID,
		Type:        t,
		TypeVersion: tv,
		State:       EventStateCreated.String(),
		CreatedAt:   time.Now().UTC(),
		ScheduledAt: scheduledAt,
		Data:        nil,
		Retry:       0,
		MaxRetry:    maxRetry,
	}
}

func (e *BaseEvent) GetID() uuid.UUID {
	return e.ID
}

func (e *BaseEvent) GetContextID() uuid.UUID {
	return e.ContextID
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

func (e *BaseEvent) GetCompletedAt() time.Time {
	return e.CompletedAt
}

func (e *BaseEvent) GetEventData() []byte {
	return e.Data
}

func (e *BaseEvent) GetScheduledAt() time.Time {
	return e.ScheduledAt
}

func (e *BaseEvent) GetRetry() int {
	return e.Retry
}

func (e *BaseEvent) GetMaxRetry() int {
	return e.MaxRetry
}

func (e *BaseEvent) Schedule(t time.Time) {
	e.ScheduledAt = t
}
