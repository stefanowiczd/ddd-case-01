package account

import "time"

//go:generate mockgen -destination=./mock/account_mock.go -package=mock -source=./account_interface.go

// Event represents a domain event
type Event interface {
	// GetID returns the unique identifier of the event
	GetID() string

	// GetType returns the type of the event
	GetType() string

	// GetAggregateID returns the unique identifier of the aggregate that the event belongs to
	GetAggregateID() string

	// GetCreatedAt returns the date and time the event was created
	GetCreatedAt() time.Time
}
