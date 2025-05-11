package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	eventdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stefanowiczd/ddd-case-01/orchestrator/infra/repo/query"
)

// FindAllEvents finds all events
func (r *OrchestratorRepository) FindAllEvents(ctx context.Context) ([]*eventdomain.BaseEvent, error) {
	events, err := r.Q.FindEvents(ctx)
	if err != nil {
		return nil, err
	}

	orchestratorEvents := make([]*eventdomain.BaseEvent, len(events))
	for i, ev := range events {
		orchestratorEvents[i] = &eventdomain.BaseEvent{
			ID:          ev.ID.Bytes,
			ContextID:   ev.ContextID.Bytes,
			Origin:      ev.EventOrigin,
			Type:        ev.EventType,
			TypeVersion: ev.EventTypeVersion,
			State:       ev.EventState,
			CreatedAt:   ev.CreatedAt.Time,
			ScheduledAt: ev.ScheduledAt.Time,
			StartedAt:   ev.StartedAt.Time,
			CompletedAt: ev.CompletedAt.Time,
			Retry:       int(ev.Retry),
			MaxRetry:    int(ev.MaxRetry),
			Data:        ev.EventData,
		}
	}

	return orchestratorEvents, nil
}

// FindProcessableEvents finds all events that are ready to be processed
func (r *OrchestratorRepository) FindProcessableEvents(ctx context.Context, limit int) ([]*eventdomain.BaseEvent, error) {
	ev, err := r.Q.FindProcessableEvents(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	orchestratorEvents := make([]*eventdomain.BaseEvent, len(ev))
	for i, ev := range ev {
		orchestratorEvents[i] = &eventdomain.BaseEvent{
			ID:          ev.ID.Bytes,
			ContextID:   ev.ContextID.Bytes,
			Origin:      ev.EventOrigin,
			Type:        ev.EventType,
			TypeVersion: ev.EventTypeVersion,
			State:       ev.EventState,
			CreatedAt:   ev.CreatedAt.Time,
			ScheduledAt: ev.ScheduledAt.Time,
			StartedAt:   ev.StartedAt.Time,
			CompletedAt: ev.CompletedAt.Time,
			Retry:       int(ev.Retry),
			MaxRetry:    int(ev.MaxRetry),
			Data:        ev.EventData,
		}
	}
	return orchestratorEvents, nil
}

// FindEventsByOriginAndStatus finds all events by origin and status
func (r *OrchestratorRepository) findByOriginAndStatus(ctx context.Context, origin, state string, limit int) ([]*eventdomain.BaseEvent, error) {
	ev, err := r.Q.FindEventsByOriginAndStatus(ctx, query.FindEventsByOriginAndStatusParams{
		EventOrigin: origin,
		EventState:  state,
		Limit:       int32(limit),
	})
	if err != nil {
		return nil, err
	}

	orchestratorEvents := make([]*eventdomain.BaseEvent, len(ev))
	for i, ev := range ev {
		orchestratorEvents[i] = &eventdomain.BaseEvent{
			ID:          ev.ID.Bytes,
			ContextID:   ev.ContextID.Bytes,
			Origin:      ev.EventOrigin,
			Type:        ev.EventType,
			TypeVersion: ev.EventTypeVersion,
			State:       ev.EventState,
			CreatedAt:   ev.CreatedAt.Time,
			ScheduledAt: ev.ScheduledAt.Time,
			StartedAt:   ev.StartedAt.Time,
			CompletedAt: ev.CompletedAt.Time,
			Retry:       int(ev.Retry),
			MaxRetry:    int(ev.MaxRetry),
			Data:        ev.EventData,
		}
	}

	return orchestratorEvents, nil
}

// FindEventByID finds an event by id
func (r *OrchestratorRepository) findByID(ctx context.Context, id uuid.UUID) (*eventdomain.BaseEvent, error) {
	ev, err := r.Q.FindEventByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return &eventdomain.BaseEvent{}, err
	}

	return &eventdomain.BaseEvent{
		ID:          ev.ID.Bytes,
		ContextID:   ev.ContextID.Bytes,
		Origin:      ev.EventOrigin,
		Type:        ev.EventType,
		TypeVersion: ev.EventTypeVersion,
		State:       ev.EventState,
		CreatedAt:   ev.CreatedAt.Time,
		ScheduledAt: ev.ScheduledAt.Time,
		StartedAt:   ev.StartedAt.Time,
		CompletedAt: ev.CompletedAt.Time,
		Retry:       int(ev.Retry),
		MaxRetry:    int(ev.MaxRetry),
		Data:        ev.EventData,
	}, nil
}
