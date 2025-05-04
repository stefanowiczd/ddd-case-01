package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	"github.com/stefanowiczd/ddd-case-01/internal/infra/repo/query"
)

// AccountRepository is a repository for account operations
type AccountRepository struct {
	Conn *pgxpool.Pool
	Q    *query.Queries
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(c *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		Conn: c,
		Q:    query.New(c),
	}
}

// CreateAccount creates a new account
func (r *AccountRepository) CreateAccount(ctx context.Context, acc *accountdomain.Account) (*accountdomain.Account, error) {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: create account: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	account, err := r.Q.CreateAccount(
		ctx,
		query.CreateAccountParams{
			ID:            pgtype.UUID{Bytes: acc.ID, Valid: true},
			CustomerID:    pgtype.UUID{Bytes: acc.CustomerID, Valid: true},
			AccountNumber: acc.AccountNumber,
			Balance:       acc.Balance,
			Currency:      acc.Currency,
			Status:        acc.Status.String(),
			CreatedAt:     pgtype.Timestamp{Time: acc.CreatedAt, Valid: true},
			UpdatedAt:     pgtype.Timestamp{Time: acc.UpdatedAt, Valid: true},
		})
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr):
			switch pgErr.Code {
			case query.POSTGRESQL_DUPLICATE_KEY_CODE:
				return nil, fmt.Errorf("checking db error: executing query: create account: %w", accountdomain.ErrAccountAlreadyExists)
			default:
				return nil, fmt.Errorf("checking db error: executing query: create account: %w", err)
			}
		}
		return nil, fmt.Errorf("executing query: create account: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("committing transaction: create account: %w", err)
	}

	return &accountdomain.Account{
		ID:            account.ID.Bytes,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Currency:      account.Currency,
		Status:        accountdomain.AccountStatus(account.Status),
	}, nil
}

// FindByID retrieves an account by its ID
func (r *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error) {
	account, err := r.Q.FindAccountByID(
		ctx,
		pgtype.UUID{Bytes: id, Valid: true},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("finding account by id: %w", accountdomain.ErrAccountNotFound)
		}

		return nil, fmt.Errorf("finding account by id: %w", err)

	}

	return &accountdomain.Account{
		ID:            account.ID.Bytes,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Currency:      account.Currency,
		Status:        accountdomain.AccountStatus(account.Status),
		CreatedAt:     account.CreatedAt.Time,
		UpdatedAt:     account.UpdatedAt.Time,
	}, nil
}

// FindByCustomerID retrieves all accounts by a customer ID
func (r *AccountRepository) FindByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*accountdomain.Account, error) {
	accounts, err := r.Q.FindAccountsByCustomerID(
		ctx,
		pgtype.UUID{Bytes: customerID, Valid: true},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*accountdomain.Account{}, nil
		}

		return nil, fmt.Errorf("finding accounts by customer id: %w", err)
	}

	accountsDomain := make([]*accountdomain.Account, 0)
	for _, account := range accounts {
		accountsDomain = append(accountsDomain, &accountdomain.Account{
			ID:            account.ID.Bytes,
			AccountNumber: account.AccountNumber,
			Balance:       account.Balance,
			Currency:      account.Currency,
			Status:        accountdomain.AccountStatus(account.Status),
			CreatedAt:     account.CreatedAt.Time,
			UpdatedAt:     account.UpdatedAt.Time,
		})
	}
	return accountsDomain, nil
}
