//go:build integration

package customer

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stretchr/testify/require"
)

func TestCustomerEventRepository_CreateCustomerEvent(t *testing.T) {

	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewCustomerEventRepository(pool)

	id := uuid.MustParse("00000000-0000-0000-0000-111111111111")

	event := &customerdomain.CustomerCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          id,
			ContextID:   uuid.MustParse("00000000-0000-0000-0000-111111111111"),
			Origin:      customerdomain.EventOrigin("customer").String(),
			Type:        customerdomain.CustomerCreatedEventType.String(),
			TypeVersion: "1.0.0",
			State:       event.EventStateCreated.String(),
			CreatedAt:   time.Now().UTC(),
			ScheduledAt: time.Now().UTC(),
			Retry:       0,
			MaxRetry:    3,
		},
		FirstName:   "John",
		LastName:    "Doe",
		Phone:       "1234567890",
		Email:       "john.doe.3@example.com",
		DateOfBirth: "1990-01-01",
		Address: customerdomain.Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
	}

	data, _ := json.Marshal(event)
	event.Data = data

	ev, err := repo.FindCustomerEventByID(ctx, id)
	require.ErrorIs(t, err, customerdomain.ErrCustomerEventNotFound)

	idOut, err := repo.CreateCustomerEvent(
		ctx,
		event,
	)
	require.NoError(t, err)
	require.Equal(t, id, idOut)

	ev, err = repo.FindCustomerEventByID(ctx, id)
	require.NotNil(t, ev)

	restoredEvent := &customerdomain.CustomerCreatedEvent{}
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

func TestCustomerEventRepository_FindCustomerEventByID(t *testing.T) {
}
