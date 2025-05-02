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
	id := uuid.New()
	now := time.Now().UTC()

	event := &AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			AccountID:   id,
			AggregateID: id,
			Type:        string(AccountCreatedEventType),
			TypeVersion: "0.0.0",
			State:       string(event.EventStateCreated),
			CreatedAt:   now,
			CompletedAt: time.Time{},
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
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
	require.Equal(t, event.GetAccountID(), restoredEvent.GetAccountID())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())

	require.Equal(t, event.InitialBalance, restoredEvent.InitialBalance)
}

func Test_AccountBlockedEvent(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	event := &AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			AccountID:   id,
			AggregateID: id,
			Type:        AccountBlockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateCreated.String(),
			CreatedAt:   now,
			CompletedAt: time.Time{},
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := AccountBlockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetAccountID(), restoredEvent.GetAccountID())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())
}

func Test_AccountUnblockedEvent(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	event := &AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			AccountID:   id,
			AggregateID: id,
			Type:        AccountUnblockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateCreated.String(),
			CreatedAt:   now,
			CompletedAt: time.Time{},
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	data, err := json.Marshal(event)
	require.NoError(t, err)

	event.Data = data

	restoredEvent := AccountUnblockedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetAccountID(), restoredEvent.GetAccountID())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())
}

func Test_FundsWithdrawnEvent(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	event := &FundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			AccountID:   id,
			AggregateID: id,
			Type:        AccountFundsWithdrawnEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateCreated.String(),
			CreatedAt:   now,
			CompletedAt: time.Time{},
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

	restoredEvent := FundsWithdrawnEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetAccountID(), restoredEvent.GetAccountID())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())

	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}

func Test_FundsDepositedEvent(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()

	event := &FundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			AccountID:   id,
			AggregateID: id,
			Type:        AccountFundsDepositedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateCreated.String(),
			CreatedAt:   now,
			CompletedAt: time.Time{},
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

	restoredEvent := FundsDepositedEvent{}
	require.NoError(t, json.Unmarshal(event.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetAccountID(), restoredEvent.GetAccountID())
	require.Equal(t, event.GetAggregateID(), restoredEvent.GetAggregateID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())

	require.Equal(t, event.Amount, restoredEvent.Amount)
	require.Equal(t, event.Balance, restoredEvent.Balance)
	require.Equal(t, event.Currency, restoredEvent.Currency)
}
