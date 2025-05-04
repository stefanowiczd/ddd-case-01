//go:build unit

package customer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

func compareCustomerBaseEvents(t *testing.T, event, restoredEvent Event) {
	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
	require.Equal(t, event.GetOrigin(), restoredEvent.GetOrigin())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())
}

func Test_CustomerCreatedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()
	origin := CustomerOrigin("customer")

	event := &CustomerCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Origin:      origin.String(),
			Type:        string(CustomerCreatedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: "1990-01-01",
		Address: Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerCreatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.FirstName, restoredEvent.FirstName)
	require.Equal(t, event.LastName, restoredEvent.LastName)
	require.Equal(t, event.Phone, restoredEvent.Phone)
	require.Equal(t, event.Email, restoredEvent.Email)
	require.Equal(t, event.DateOfBirth, restoredEvent.DateOfBirth)
	require.Equal(t, event.Address, restoredEvent.Address)
}

func Test_CustomerActivatedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerActivatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        string(CustomerActivatedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerActivatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

}

func Test_CustomerDeactivatedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerDeactivatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        string(CustomerDeactivatedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerDeactivatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)
}

func Test_CustomerBlockedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()
	reason := "Blocked by customer"

	event := &CustomerBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        string(CustomerBlockedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
		Reason: reason,
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerBlockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.Reason, restoredEvent.Reason)
}

func Test_CustomerUnblockedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        string(CustomerUnblockedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerUnblockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)
}

func Test_CustomerUpdatedEvent_Name(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerUpdatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        CustomerUpdatedNameEventType.String(),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: "1990-01-01",
		Address: Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerUpdatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.FirstName, restoredEvent.FirstName)
	require.Equal(t, event.LastName, restoredEvent.LastName)
	require.Equal(t, event.Phone, restoredEvent.Phone)
	require.Equal(t, event.Email, restoredEvent.Email)
	require.Equal(t, event.DateOfBirth, restoredEvent.DateOfBirth)
	require.Equal(t, event.Address, restoredEvent.Address)
}

func Test_CustomerUpdatedEvent_Contact(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerUpdatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        CustomerUpdatedContactEventType.String(),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: "1990-01-01",
		Address: Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerUpdatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.FirstName, restoredEvent.FirstName)
	require.Equal(t, event.LastName, restoredEvent.LastName)
	require.Equal(t, event.Phone, restoredEvent.Phone)
	require.Equal(t, event.Email, restoredEvent.Email)
	require.Equal(t, event.DateOfBirth, restoredEvent.DateOfBirth)
	require.Equal(t, event.Address, restoredEvent.Address)
}

func Test_CustomerUpdatedEvent_Address(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerUpdatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        CustomerUpdatedAddressEventType.String(),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: "1990-01-01",
		Address: Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerUpdatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.FirstName, restoredEvent.FirstName)
	require.Equal(t, event.LastName, restoredEvent.LastName)
	require.Equal(t, event.Phone, restoredEvent.Phone)
	require.Equal(t, event.Email, restoredEvent.Email)
	require.Equal(t, event.Address, restoredEvent.Address)
}

func Test_CustomerUpdatedEvent_All(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerUpdatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        CustomerUpdatedAllEventType.String(),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe@example.com",
		DateOfBirth: "1990-01-01",
		Address: Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerUpdatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.FirstName, restoredEvent.FirstName)
	require.Equal(t, event.LastName, restoredEvent.LastName)
	require.Equal(t, event.Phone, restoredEvent.Phone)
	require.Equal(t, event.Email, restoredEvent.Email)
	require.Equal(t, event.DateOfBirth, restoredEvent.DateOfBirth)
	require.Equal(t, event.Address, restoredEvent.Address)
}

func Test_CustomerDeletedEvent(t *testing.T) {
	now := time.Now().UTC()
	eventID := uuid.New()
	customerID := uuid.New()

	event := &CustomerDeletedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   customerID,
			Type:        CustomerDeletedEventType.String(),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &CustomerDeletedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)
}
