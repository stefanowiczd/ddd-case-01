package customer

import (
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

// CustomerEventType represents the type of customer event
type CustomerEventType string

// String returns the string representation of the customer event type
func (e CustomerEventType) String() string {
	return string(e)
}

const (
	CustomerCreatedEventType CustomerEventType = "customer.created"
	CustomerDeletedEventType CustomerEventType = "customer.deleted"

	CustomerActivatedEventType   CustomerEventType = "customer.activated"
	CustomerDeactivatedEventType CustomerEventType = "customer.deactivated"

	CustomerBlockedEventType   CustomerEventType = "customer.blocked"
	CustomerUnblockedEventType CustomerEventType = "customer.unblocked"

	CustomerUpdatedNameEventType    CustomerEventType = "customer.updated.name"
	CustomerUpdatedContactEventType CustomerEventType = "customer.updated.contact"
	CustomerUpdatedAddressEventType CustomerEventType = "customer.updated.address"
	CustomerUpdatedAllEventType     CustomerEventType = "customer.updated.all"
)

// AccountCreatedEvent is emitted when a new account is created
type CustomerCreatedEvent struct {
	event.BaseEvent
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

// CustomerActivatedEvent is emitted when a customer is activated
type CustomerActivatedEvent struct {
	event.BaseEvent
}

// CustomerDeactivatedEvent is emitted when a customer is deactivated
type CustomerDeactivatedEvent struct {
	event.BaseEvent
}

// CustomerBlockedEvent is emitted when a customer is blocked
type CustomerBlockedEvent struct {
	event.BaseEvent
	Reason string `json:"reason"` // The reason the customer was blocked
}

// CustomerUnblockedEvent is emitted when a customer is unblocked
type CustomerUnblockedEvent struct {
	event.BaseEvent
}

// CustomerUpdatedEvent is emitted when a customer is updated
type CustomerUpdatedEvent struct {
	event.BaseEvent
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

// CustomerDeletedEvent is emitted when a customer is deleted
type CustomerDeletedEvent struct {
	event.BaseEvent
}
