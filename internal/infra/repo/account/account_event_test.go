//go:build integration

package account

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

func TestAccountEventRepository_CreateAccountEvent(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountEventRepository(pool)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	event := accountdomain.AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.MustParse("00000000-1111-2222-0000-000000000000"),
			ContextID:   id,
			Origin:      accountdomain.AccountOrigin("account").String(),
			Type:        accountdomain.AccountCreatedEventType.String(),
			TypeVersion: "0.0.1",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		InitialBalance: 10,
	}

	data, _ := json.Marshal(event)
	event.Data = data

	ev, err := repo.FindAccountEventByID(ctx, id)
	require.ErrorIs(t, err, accountdomain.ErrAccountEventNotFound)

	ID, err := repo.CreateAccountEvent(ctx, &event)
	require.NoError(t, err)

	ev, err = repo.FindAccountEventByID(ctx, ID)
	require.NoError(t, err)
	require.NotNil(t, ev)
	require.Equal(t, event.GetType(), ev.GetType())

	restoredEvent := accountdomain.AccountCreatedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
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

func TestAccountEventRepository_AccountBlockEvent(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountEventRepository(pool)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	event := accountdomain.AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.MustParse("00000000-1111-2222-0000-000000000000"),
			ContextID:   id,
			Type:        accountdomain.AccountBlockedEventType.String(),
			TypeVersion: "0.0.1",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	ev, err := repo.FindAccountEventByID(ctx, id)
	require.ErrorIs(t, err, accountdomain.ErrAccountEventNotFound)

	data, _ := json.Marshal(event)
	event.Data = data

	ID, err := repo.CreateAccountEvent(ctx, &event)
	require.NoError(t, err)

	ev, err = repo.FindAccountEventByID(ctx, ID)
	require.NoError(t, err)
	require.NotNil(t, ev)

	restoredEvent := accountdomain.AccountBlockedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())
}

func TestAccountEventRepository_AccountUnblockEvent(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountEventRepository(pool)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	event := accountdomain.AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.MustParse("00000000-1111-2222-0000-000000000000"),
			ContextID:   id,
			Type:        accountdomain.AccountUnblockedEventType.String(),
			TypeVersion: "0.0.1",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	}

	ev, err := repo.FindAccountEventByID(ctx, id)
	require.ErrorIs(t, err, accountdomain.ErrAccountEventNotFound)

	data, _ := json.Marshal(event)
	event.Data = data

	ID, err := repo.CreateAccountEvent(ctx, &event)
	require.NoError(t, err)

	ev, err = repo.FindAccountEventByID(ctx, ID)
	require.NoError(t, err)
	require.NotNil(t, ev)

	restoredEvent := accountdomain.AccountUnblockedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))
	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
	require.Equal(t, event.GetType(), restoredEvent.GetType())
	require.Equal(t, event.GetTypeVersion(), restoredEvent.GetTypeVersion())
	require.Equal(t, event.GetState(), restoredEvent.GetState())
	require.Equal(t, event.GetCreatedAt(), restoredEvent.GetCreatedAt())
	require.Equal(t, event.GetCompletedAt(), restoredEvent.GetCompletedAt())
	require.Equal(t, event.GetScheduledAt(), restoredEvent.GetScheduledAt())
	require.Equal(t, event.GetRetry(), restoredEvent.GetRetry())
	require.Equal(t, event.GetMaxRetry(), restoredEvent.GetMaxRetry())
}

func TestAccountEventRepository_FundsWithdrawnEvent(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountEventRepository(pool)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	event := accountdomain.FundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.MustParse("00000000-1111-2222-0000-000000000000"),
			ContextID:   id,
			Type:        accountdomain.AccountFundsWithdrawnEventType.String(),
			TypeVersion: "0.0.1",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	ev, err := repo.FindAccountEventByID(ctx, id)
	require.ErrorIs(t, err, accountdomain.ErrAccountEventNotFound)

	data, _ := json.Marshal(event)
	event.Data = data

	ID, err := repo.CreateAccountEvent(ctx, &event)
	require.NoError(t, err)

	ev, err = repo.FindAccountEventByID(ctx, ID)
	require.NoError(t, err)
	require.NotNil(t, ev)

	restoredEvent := accountdomain.FundsWithdrawnEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))
	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
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

func TestAccountEventRepository_FundsDepositedEvent(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountEventRepository(pool)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	event := accountdomain.FundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.MustParse("00000000-1111-2222-0000-000000000000"),
			ContextID:   id,
			Type:        accountdomain.AccountFundsDepositedEventType.String(),
			TypeVersion: "0.0.1",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			CompletedAt: time.Time{},
			ScheduledAt: time.Now().UTC(),
			Retry:       1,
			MaxRetry:    3,
			Data:        nil,
		},
		Amount:   10,
		Balance:  10,
		Currency: "USD",
	}

	ev, err := repo.FindAccountEventByID(ctx, id)
	require.ErrorIs(t, err, accountdomain.ErrAccountEventNotFound)

	data, _ := json.Marshal(event)
	event.Data = data

	ID, err := repo.CreateAccountEvent(ctx, &event)
	require.NoError(t, err)

	ev, err = repo.FindAccountEventByID(ctx, ID)
	require.NoError(t, err)
	require.NotNil(t, ev)

	restoredEvent := accountdomain.FundsDepositedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	require.Equal(t, event.GetID(), restoredEvent.GetID())
	require.Equal(t, event.GetContextID(), restoredEvent.GetContextID())
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
