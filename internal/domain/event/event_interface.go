package event

// EventPublisher defines the interface for publishing domain events
type EventPublisher interface {
	Publish(event Event) error
}

// EventStore defines the interface for storing domain events
type EventStore interface {
	Save(event Event) error
	Load(aggregateID string) ([]Event, error)
}
