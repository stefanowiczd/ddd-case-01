package account

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	query "github.com/stefanowiczd/ddd-case-01/internal/infra/repo/query"
)

const (
	// POSTGRESQL_DUPLICATE_KEY_CODE is the code for a duplicate key value violation
	POSTGRESQL_DUPLICATE_KEY_CODE = "23505"
)

var (
	// ErrNoRows is returned when no rows are found in the result set
	ErrNoRows = errors.New("no rows in result set")
	// ErrDuplicateKey is returned when a duplicate key value violates a unique constraint definition
	ErrDuplicateKey = errors.New("duplicate key value violates unique constraint definition")
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

// FindByID retrieves an account by its ID
func (r *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error) {
	account, err := r.Q.FindAccountByID(
		ctx,
		pgtype.UUID{Bytes: id, Valid: true},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, accountdomain.ErrAccountNotFound
		}

		return nil, fmt.Errorf("finding account by id: %w", err)

	}

	return &accountdomain.Account{
		ID:            account.ID.String(),
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Currency:      account.Currency,
		Status:        accountdomain.AccountStatus(account.Status),
		CreatedAt:     account.CreatedAt.Time,
		UpdatedAt:     account.UpdatedAt.Time,
	}, nil
}

// FindByAccountNumber retrieves an account by its account number
func (r *AccountRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (*accountdomain.Account, error) {
	account, err := r.Q.FindAccountByNumber(
		ctx,
		accountNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("finding account by account number: %w", ErrNoRows)
		}

		return nil, fmt.Errorf("finding account by account number: %w", err)
	}

	return &accountdomain.Account{
		ID:            account.ID.String(),
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
		return nil, fmt.Errorf("finding accounts by customer id: %w", err)
	}

	accountsDomain := make([]*accountdomain.Account, 0)
	for _, account := range accounts {
		accountsDomain = append(accountsDomain, &accountdomain.Account{
			ID:            account.ID.String(),
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
			CustomerID:    pgtype.UUID{Bytes: uuid.MustParse(acc.CustomerID), Valid: true},
			AccountNumber: acc.AccountNumber,
			Balance:       acc.Balance,
			Currency:      acc.Currency,
			Status:        string(acc.Status),
		})
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr):
			switch pgErr.Code {
			case POSTGRESQL_DUPLICATE_KEY_CODE:
				return nil, fmt.Errorf("checking db error: executing query: create account: %w", ErrDuplicateKey)
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
		ID:            account.ID.String(),
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Currency:      account.Currency,
		Status:        accountdomain.AccountStatus(account.Status),
	}, nil
}
