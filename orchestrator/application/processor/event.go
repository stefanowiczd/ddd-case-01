package processor

import (
	"encoding/json"
)

// Event represents the base structure for all events
type Event[T any] struct {
	Data T
	Ok   bool
}

// UnmarshalEvent unmarshals the event data into the generic type
func UnmarshalEvent[T any](eventData []byte) (Event[T], error) {
	if len(eventData) == 0 {
		return Event[T]{}, nil
	}

	var data T
	if err := json.Unmarshal(eventData, &data); err != nil {
		return Event[T]{}, err
	}

	return Event[T]{
		Data: data,
		Ok:   true,
	}, nil
}
