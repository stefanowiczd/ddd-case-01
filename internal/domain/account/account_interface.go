package account

import "time"

//go:generate mockgen -destination=./mock/account_mock.go -package=account -source=./account_interface.go

// Event represents a domain event
type Event interface {
	GetID() string
	GetType() string
	GetAggregateID() string
	GetCreatedAt() time.Time
}
