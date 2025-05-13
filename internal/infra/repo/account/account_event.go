package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
	"github.com/stefanowiczd/ddd-case-01/internal/infra/repo/query"
)

// AccountEventRepository is a repository for account event operations
type AccountEventRepository struct {
	Conn *pgxpool.Pool
	Q    *query.Queries
}

// NewAccountEventRepository creates a new account event repository
func NewAccountEventRepository(c *pgxpool.Pool) *AccountEventRepository {
	return &AccountEventRepository{
		Conn: c,
		Q:    query.New(c),
	}
}

// CreateAccountEvent creates an account event
func (r *AccountEventRepository) CreateAccountEvent(ctx context.Context, eventObject BaseEvent) (uuid.UUID, error) {
	switch eventObject.(type) {
	case *accountdomain.AccountCreatedEvent,
		*accountdomain.AccountBlockedEvent,
		*accountdomain.AccountUnblockedEvent,
		*accountdomain.AccountFundsWithdrawnEvent,
		*accountdomain.AccountFundsDepositedEvent:

		tx, err := r.Conn.Begin(ctx)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("starting transaction: creating account event: %w", err)
		}
		defer func() { _ = tx.Rollback(ctx) }()

		qtx := r.Q.WithTx(tx)

		accountEvent, err := qtx.CreateAccountEvent(
			ctx,
			query.CreateAccountEventParams{
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
			return uuid.Nil, fmt.Errorf("creating account event: %w", err)
		}

		if err = tx.Commit(ctx); err != nil {
			return uuid.UUID{}, fmt.Errorf("committing transaction: creating account event: %w", err)
		}
		return accountEvent.ID.Bytes, nil
	default:
		return uuid.Nil, fmt.Errorf("casting account event: event type not recognized")
	}
}

// FindAccountEventByID finds an account event by id
func (r *AccountEventRepository) FindAccountEventByID(ctx context.Context, id uuid.UUID) (BaseEvent, error) {
	ev, err := r.Q.FindAccountEventByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &event.BaseEvent{}, fmt.Errorf("finding account event by id: %w", accountdomain.ErrAccountEventNotFound)
		}

		return &event.BaseEvent{}, fmt.Errorf("finding account event by id: %w", err)
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
