package customer

import (
	"time"

	"github.com/google/uuid"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

type Customer struct {
	ID        uuid.UUID      `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     string         // Phone number of the customer
	Email     string         // Email address of the customer
	Address   Address        // Physical address of the customer
	Status    CustomerStatus // Current status of the customer
	Accounts  []string       // List of account IDs associated with the customer
	CreatedAt time.Time      // When the customer was created
	UpdatedAt time.Time      // When the customer was last updated
	Events    []Event        // List of events associated with the customer
}

// Address represents a physical address
type Address struct {
	Street     string `json:"street"`      // Street name and number
	City       string `json:"city"`        // City name
	State      string `json:"state"`       // State or province
	PostalCode string `json:"postal_code"` // Postal or ZIP code
	Country    string `json:"country"`     // Country name
}

func (a Address) compare(b Address) bool { //nolint:unused
	return a.Street == b.Street &&
		a.City == b.City &&
		a.State == b.State &&
		a.PostalCode == b.PostalCode &&
		a.Country == b.Country
}

// CustomerStatus represents the status of a customer
type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
	CustomerStatusBlocked  CustomerStatus = "blocked"
)

func (a CustomerStatus) String() string {
	return string(a)
}

// NewCustomer creates a new customer
func NewCustomer(id uuid.UUID, firstName string, lastName string, phone string, email string, address Address) *Customer {
	now := time.Now().UTC()

	customer := &Customer{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Address:   address,
		Status:    CustomerStatusActive, // TODO: change to inactive
		Accounts:  []string{},
		CreatedAt: now,
		UpdatedAt: now,
		Events:    []Event{},
	}

	customer.addEvent(
		&CustomerCreatedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   id,
				Type:        CustomerCreatedEventType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				CompletedAt: time.Time{},
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
				Data:        nil,
			},
			FirstName: firstName,
			LastName:  lastName,
			Phone:     phone,
			Email:     email,
			Address:   address,
		})

	return customer
}

// Activate activates a customer
func (c *Customer) Activate() {
	now := time.Now().UTC()
	c.Status = CustomerStatusActive
	c.UpdatedAt = now

	c.Events = append(
		c.Events,
		&CustomerActivatedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   c.ID,
				Type:        CustomerActivatedEventType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
			},
		})
}

// Deactivate deactivates a customer
func (c *Customer) Deactivate() {
	now := time.Now().UTC()
	c.Status = CustomerStatusInactive
	c.UpdatedAt = now

	c.Events = append(
		c.Events,
		&CustomerDeactivatedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   c.ID,
				Type:        CustomerDeactivatedEventType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
			},
		})
}

// Block blocks a customer
func (c *Customer) Block(reason string) {
	now := time.Now().UTC()
	c.Status = CustomerStatusBlocked
	c.UpdatedAt = now

	c.Events = append(
		c.Events,
		&CustomerBlockedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   c.ID,
				Type:        CustomerBlockedEventType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
			},
			Reason: reason,
		})
}

// Unblock unblocks a customer
func (c *Customer) Unblock() {
	now := time.Now().UTC()
	c.Status = CustomerStatusActive
	c.UpdatedAt = now

	c.Events = append(
		c.Events,
		&CustomerUnblockedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   c.ID,
				Type:        CustomerUnblockedEventType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
			},
		})
}

func (c *Customer) Update(
	updateType CustomerEventType,
	firstName, lastName string,
	phone, email string,
	address Address,
) {
	now := time.Now().UTC()
	c.UpdatedAt = now

	c.FirstName = firstName
	c.LastName = lastName
	c.Phone = phone
	c.Email = email
	c.Address = address

	c.Events = append(
		c.Events,
		&CustomerUpdatedEvent{
			BaseEvent: event.BaseEvent{
				ID:          uuid.New(),
				ContextID:   c.ID,
				Type:        updateType.String(),
				TypeVersion: "0.0.0",
				State:       event.EventStateCreated.String(),
				CreatedAt:   now,
				ScheduledAt: now,
				Retry:       0,
				MaxRetry:    3,
			},
			FirstName: firstName,
			LastName:  lastName,
			Phone:     phone,
			Email:     email,
			Address:   address,
		})
}

// GetEvents returns all domain events that have occurred on this customer.
func (c *Customer) GetEvents() []Event {
	return c.Events
}

// ClearEvents removes all recorded events from the customer.
// This is typically called after events have been processed.
func (c *Customer) ClearEvents() {
	c.Events = make([]Event, 0)
}

// addEvent is an internal method to record a new domain event.
func (c *Customer) addEvent(event Event) {
	c.Events = append(c.Events, event)
}
