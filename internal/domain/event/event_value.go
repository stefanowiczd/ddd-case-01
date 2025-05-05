package event

// EventOrigin represents the origin of an event, i.e. account, customer, etc.
type EventOrigin string

func (e EventOrigin) String() string {
	return string(e)
}

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
