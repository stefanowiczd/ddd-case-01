//go:build unit

package account

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stretchr/testify/require"
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

func Test_AccountCreatedEvent(t *testing.T) {
	eventID := uuid.New()
	accountID := uuid.New()
	customerID := uuid.New()
	now := time.Now().UTC()
	origin := EventOrigin("account")

	event := &AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   accountID,
			Origin:      origin.String(),
			Type:        AccountCreatedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		CustomerID:     customerID,
		InitialBalance: 10,
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &AccountCreatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.InitialBalance, restoredEvent.InitialBalance)
	require.Equal(t, event.CustomerID, restoredEvent.CustomerID)
}

func Test_AccountBlockedEvent(t *testing.T) {
	eventID := uuid.New()
	accountID := uuid.New()
	now := time.Now().UTC()
	origin := EventOrigin("account")

	event := &AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   accountID,
			Origin:      origin.String(),
			Type:        AccountBlockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &AccountBlockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)
}

func Test_AccountUnblockedEvent(t *testing.T) {
	eventID := uuid.New()
	accountID := uuid.New()
	now := time.Now().UTC()
	origin := EventOrigin("account")

	event := &AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   accountID,
			Origin:      origin.String(),
			Type:        AccountUnblockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &AccountUnblockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)
}

func Test_FundsWithdrawnEvent(t *testing.T) {
	eventID := uuid.New()
	accountID := uuid.New()
	now := time.Now().UTC()
	origin := EventOrigin("account")

	event := &AccountFundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   accountID,
			Origin:      origin.String(),
			Type:        AccountFundsWithdrawnEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &AccountFundsWithdrawnEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}

func Test_FundsDepositedEvent(t *testing.T) {
	eventID := uuid.New()
	accountID := uuid.New()
	now := time.Now().UTC()
	origin := EventOrigin("account")

	event := &AccountFundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          eventID,
			ContextID:   accountID,
			Origin:      origin.String(),
			Type:        AccountFundsDepositedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := &AccountFundsDepositedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	compareCustomerBaseEvents(t, event, restoredEvent)

	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}
