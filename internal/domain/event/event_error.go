package event

import (
	"errors"
)

// Event errors
var (
	// ErrEventNotFound is returned when an event is not found
	ErrEventNotFound = errors.New("event not found") // TODO: decide if we should use this error for all domains?
)
