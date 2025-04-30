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

func Test_AccountCreatedEvent(t *testing.T) {
	event := &AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			Type:        AccountCreatedEventType,
			TypeVersion: "0.0.0",
			AggregateID: uuid.New(),
			CreatedAt:   time.Now().UTC(),
			Data:        nil,
		},
		InitialBalance: 10,
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := AccountCreatedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.InitialBalance, restoredEvent.InitialBalance)
}

func Test_AccountBlockedEvent(t *testing.T) {
	event := &AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			Type:        AccountBlockedEventType,
			TypeVersion: "0.0.0",
			AggregateID: uuid.New(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := AccountBlockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
}

func Test_AccountUnblockedEvent(t *testing.T) {
	event := &AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			Type:        AccountUnblockedEventType,
			TypeVersion: "0.0.0",
			AggregateID: uuid.New(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := AccountUnblockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
}

func Test_FundsWithdrawnEvent(t *testing.T) {
	event := &FundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			Type:        AccountFundsWithdrawnEventType,
			TypeVersion: "0.0.0",
			AggregateID: uuid.New(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := FundsWithdrawnEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}

func Test_FundsDepositedEvent(t *testing.T) {
	event := &FundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			Type:        AccountFundsDepositedEventType,
			TypeVersion: "0.0.0",
			AggregateID: uuid.New(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := FundsDepositedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}
