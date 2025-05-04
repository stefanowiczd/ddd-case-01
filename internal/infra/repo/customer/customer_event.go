package customer

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stefanowiczd/ddd-case-01/internal/infra/repo/query"
)

// CustomerEventRepository is a repository for customer event operations
type CustomerEventRepository struct {
	Conn *pgxpool.Pool
	Q    *query.Queries
}

// NewCustomerEventRepository creates a new customer event repository
func NewCustomerEventRepository(conn *pgxpool.Pool) *CustomerEventRepository {
	return &CustomerEventRepository{
		Conn: conn,
		Q:    query.New(conn),
	}
}

func (r *CustomerEventRepository) CreateCustomerEvent(ctx context.Context, eventObject BaseEvent) (uuid.UUID, error) {
	switch eventObject.(type) {
	case *customerdomain.CustomerCreatedEvent,
		*customerdomain.CustomerUpdatedEvent,
		*customerdomain.CustomerDeletedEvent,
		*customerdomain.CustomerActivatedEvent,
		*customerdomain.CustomerDeactivatedEvent,
		*customerdomain.CustomerBlockedEvent,
		*customerdomain.CustomerUnblockedEvent:

		tx, err := r.Conn.Begin(ctx)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("starting transaction: creating customer event: %w", err)
		}
		defer func() { _ = tx.Rollback(ctx) }()

		qtx := r.Q.WithTx(tx)

		customerEvent, err := qtx.CreateCustomerEvent(
			ctx,
			query.CreateCustomerEventParams{
				ID:               pgtype.UUID{Bytes: eventObject.GetID(), Valid: true},
				ContextID:        pgtype.UUID{Bytes: eventObject.GetContextID(), Valid: true},
				EventOrigin:      eventObject.GetOrigin(),
				EventType:        eventObject.GetType(),
				EventTypeVersion: eventObject.GetTypeVersion(),
				EventState:       eventObject.GetState(),
				CreatedAt:        pgtype.Timestamp{Time: eventObject.GetCreatedAt(), Valid: true},
				ScheduledAt:      pgtype.Timestamp{Time: eventObject.GetScheduledAt(), Valid: true},
				Retry:            int32(eventObject.GetRetry()),
				MaxRetry:         int32(eventObject.GetMaxRetry()),
				EventData:        eventObject.GetEventData(),
			},
		)
		if err != nil {
			return uuid.Nil, fmt.Errorf("creating customer event: %w", err)
		}

		if err = tx.Commit(ctx); err != nil {
			return uuid.UUID{}, fmt.Errorf("committing transaction: creating customer event: %w", err)
		}

		return customerEvent.ID.Bytes, nil
	default:
		return uuid.Nil, fmt.Errorf("casting customer event: event type not recognized")
	}
}

func (r *CustomerEventRepository) FindCustomerEventByID(ctx context.Context, id uuid.UUID) (BaseEvent, error) {
	ev, err := r.Q.FindCustomerEventByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &event.BaseEvent{}, fmt.Errorf("finding customer event by id: %w", customerdomain.ErrCustomerEventNotFound)
		}

		return &event.BaseEvent{}, fmt.Errorf("finding customer event by id: %w", err)
	}

	return &event.BaseEvent{
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
		Retry:       int(ev.Retry),    // TODO check this one more time...
		MaxRetry:    int(ev.MaxRetry), // TODO check this one more time...
		Data:        ev.EventData,
	}, nil
}
