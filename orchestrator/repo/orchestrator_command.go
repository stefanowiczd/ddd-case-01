package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stefanowiczd/ddd-case-01/orchestrator/repo/query"
)

// UpdateEventStart updates the event at start up
func (r *OrchestratorRepository) UpdateEventStart(ctx context.Context, id uuid.UUID) error {
	err := r.Q.UpdateEventStart(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	return nil
}

// UpdateEventCompletion updates the event at completion
func (r *OrchestratorRepository) UpdateEventCompletion(ctx context.Context, id uuid.UUID) error {
	err := r.Q.UpdateEventCompletion(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}

	return nil
}

// UpdateEventRetry updates the event at retry
func (r *OrchestratorRepository) UpdateEventRetry(ctx context.Context, id uuid.UUID, retry int) error {
	err := r.Q.UpdateEventRetry(ctx, query.UpdateEventRetryParams{
		ID:            pgtype.UUID{Bytes: id, Valid: true},
		RetryInterval: retry,
	})
	if err != nil {
		return err
	}

	return nil
}
