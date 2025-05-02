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

func (r *AccountEventRepository) CreateAccountEvent(ctx context.Context, eventObject AccountEvent) (uuid.UUID, error) {

	switch eventObject.(type) {
	case *accountdomain.AccountCreatedEvent,
		*accountdomain.AccountBlockedEvent,
		*accountdomain.AccountUnblockedEvent,
		*accountdomain.FundsWithdrawnEvent,
		*accountdomain.FundsDepositedEvent:

		tx, err := r.Conn.Begin(ctx)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("starting transaction: creating account event: %w", err)
		}
		defer func() { _ = tx.Rollback(ctx) }()

		qtx := r.Q.WithTx(tx)

		accountEvent, err := qtx.CreateAccountEvent(
			ctx,
			query.CreateAccountEventParams{
				AccountID:        pgtype.UUID{Bytes: eventObject.GetAggregateID(), Valid: true},
				EventType:        eventObject.GetType(),
				EventTypeVersion: eventObject.GetTypeVersion(),
				EventState:       eventObject.GetState(),
				ScheduledAt:      pgtype.Timestamptz{Time: eventObject.GetScheduledAt(), Valid: true},
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

func (r *AccountEventRepository) FindAccountEventByID(ctx context.Context, id uuid.UUID) (query.AccountEvent, error) {
	ev, err := r.Q.FindAccountEventByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return query.AccountEvent{}, fmt.Errorf("finding account event by id: %w", ErrNoRows)
		}
		return query.AccountEvent{}, fmt.Errorf("finding account event by id: %w", err)
	}

	return ev, nil
}
