package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stefanowiczd/ddd-case-01/orchestrator/repo/query"
)

// OrchestratorRepository is the repository for the orchestrator that handles the event database operations
type OrchestratorRepository struct {
	Conn *pgxpool.Pool
	Q    *query.Queries
}

// NewOrchestratorRepository creates a new orchestrator repository
func NewOrchestratorRepository(conn *pgxpool.Pool) *OrchestratorRepository {
	return &OrchestratorRepository{
		Conn: conn,
		Q:    query.New(conn),
	}
}
