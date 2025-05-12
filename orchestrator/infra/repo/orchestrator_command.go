package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stefanowiczd/ddd-case-01/orchestrator/infra/repo/query"
)

// UpdateEventStart updates the event at start up
func (r *OrchestratorRepository) UpdateEventStart(ctx context.Context, id uuid.UUID) error {
	if err := r.Q.UpdateEventStart(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		return fmt.Errorf("updating event start: %w", err)
	}

	return nil
}

// UpdateEventCompletion updates the event at completion
func (r *OrchestratorRepository) UpdateEventCompletion(ctx context.Context, id uuid.UUID) error {
	if err := r.Q.UpdateEventCompletion(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		return fmt.Errorf("updating event completion: %w", err)
	}

	return nil
}

// UpdateEventRetry updates the event at retry
func (r *OrchestratorRepository) UpdateEventRetry(ctx context.Context, id uuid.UUID, retry int) error {
	if err := r.Q.UpdateEventRetry(ctx, query.UpdateEventRetryParams{
		ID:            pgtype.UUID{Bytes: id, Valid: true},
		RetryInterval: retry,
	}); err != nil {
		return fmt.Errorf("updating event retry: %w", err)
	}

	return nil
}
