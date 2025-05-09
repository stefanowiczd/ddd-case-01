//go:build integration

package repo

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

func TestOrchestrator_Start(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	eventRepo := NewOrchestratorRepository(pool)

	events, err := eventRepo.FindAllEvents(ctx)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 3)

	ev := events[0]

	restoredEvent := &customerdomain.CustomerCreatedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	events, err = eventRepo.findByOriginAndStatus(ctx, "customer", "ready", 5)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 2)

	ev = events[0]
	restoredEvent = &customerdomain.CustomerCreatedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	events, err = eventRepo.findByOriginAndStatus(ctx, "customer", "ready", 5)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 2)

}

func TestOrchestrator_SetStateAndScheduledAt(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	eventRepo := NewOrchestratorRepository(pool)

	events, err := eventRepo.findByOriginAndStatus(ctx, "customer", "ready", 5)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 2)

	ev := events[0]

	restoredEvent := &customerdomain.CustomerCreatedEvent{}
	require.NoError(t, json.Unmarshal(ev.GetEventData(), &restoredEvent))

	id := ev.ID

	eventBeforeUpdate, err := eventRepo.findByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, eventBeforeUpdate)
	require.Equal(t, 0, eventBeforeUpdate.Retry)

	err = eventRepo.UpdateEventRetry(ctx, eventBeforeUpdate.ID, 1)
	require.NoError(t, err)

	eventAfterUpdate, err := eventRepo.findByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, eventAfterUpdate)

	require.Equal(t, 1, eventAfterUpdate.Retry)
	require.Greater(t, eventAfterUpdate.ScheduledAt, eventBeforeUpdate.ScheduledAt)
	require.Equal(t, "ready", eventAfterUpdate.State)

	err = eventRepo.UpdateEventRetry(ctx, eventBeforeUpdate.ID, 1)
	require.NoError(t, err)

	eventAfterSecondUpdate, err := eventRepo.findByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, eventAfterSecondUpdate)

	require.Equal(t, 2, eventAfterSecondUpdate.Retry)
	require.Greater(t, eventAfterSecondUpdate.ScheduledAt, eventAfterUpdate.ScheduledAt)
	require.Equal(t, "ready", eventAfterSecondUpdate.State)

	err = eventRepo.UpdateEventRetry(ctx, eventBeforeUpdate.ID, 1)
	require.NoError(t, err)

	eventAfterThirdUpdate, err := eventRepo.findByID(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, eventAfterThirdUpdate)

	require.Equal(t, 3, eventAfterThirdUpdate.Retry)
	require.Equal(t, eventAfterThirdUpdate.ScheduledAt, eventAfterSecondUpdate.ScheduledAt)
	require.Equal(t, "failed", eventAfterThirdUpdate.State)
	require.NotNil(t, eventAfterThirdUpdate.CompletedAt)
	require.Greater(t, time.Now().UTC(), eventAfterThirdUpdate.CompletedAt.UTC())
}

func TestOrchestrator_UpdateEventCompletion(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	eventRepo := NewOrchestratorRepository(pool)

	events, err := eventRepo.findByOriginAndStatus(ctx, "customer", "ready", 5)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 2)

	ev := events[0]

	err = eventRepo.UpdateEventCompletion(ctx, ev.ID)
	require.NoError(t, err)

	eventAfterUpdate, err := eventRepo.findByID(ctx, ev.ID)
	require.NoError(t, err)
	require.NotNil(t, eventAfterUpdate)
	require.Equal(t, "completed", eventAfterUpdate.State)
	require.NotNil(t, eventAfterUpdate.CompletedAt)
}

func TestOrchestrator_UpdateEventStart(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	eventRepo := NewOrchestratorRepository(pool)

	events, err := eventRepo.findByOriginAndStatus(ctx, "customer", "ready", 5)
	require.NoError(t, err)
	require.NotNil(t, events)
	require.Len(t, events, 2)

	ev := events[0]

	err = eventRepo.UpdateEventStart(ctx, ev.ID)
	require.NoError(t, err)

	eventAfterStart, err := eventRepo.findByID(ctx, ev.ID)
	require.NoError(t, err)
	require.NotNil(t, eventAfterStart)
	require.Equal(t, "processing", eventAfterStart.State)
	require.NotNil(t, eventAfterStart.StartedAt)
	require.Greater(t, time.Now().UTC(), eventAfterStart.StartedAt.UTC())

	err = eventRepo.UpdateEventCompletion(ctx, ev.ID)
	require.NoError(t, err)

	eventAfterCompletion, err := eventRepo.findByID(ctx, ev.ID)
	require.NoError(t, err)
	require.NotNil(t, eventAfterCompletion)
	require.Equal(t, "completed", eventAfterCompletion.State)
	require.NotNil(t, eventAfterCompletion.CompletedAt)
	require.Greater(t, time.Now().UTC(), eventAfterCompletion.CompletedAt.UTC())
}
