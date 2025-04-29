package account

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Save(ctx context.Context, acc *account.Account) error {
	query := `
		INSERT INTO accounts (id, account_number, customer_id, balance, currency, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.pool.Exec(ctx, query,
		acc.ID,
		acc.AccountNumber,
		acc.CustomerID,
		acc.Balance,
		acc.Currency,
		acc.Status,
	)
	if err != nil {
		return fmt.Errorf("failed to save account: %w", err)
	}

	// Save events
	for _, event := range acc.GetEvents() {
		eventData, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %w", err)
		}

		_, err = r.pool.Exec(ctx, `
			INSERT INTO account_events (account_id, event_type, event_data)
			VALUES ($1, $2, $3)
		`, acc.ID, event.GetType(), eventData)
		if err != nil {
			return fmt.Errorf("failed to save event: %w", err)
		}
	}

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*account.Account, error) {
	query := `
		SELECT id, account_number, customer_id, balance, currency, status
		FROM accounts
		WHERE id = $1
	`

	var acc account.Account
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&acc.ID,
		&acc.AccountNumber,
		&acc.CustomerID,
		&acc.Balance,
		&acc.Currency,
		&acc.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find account by ID: %w", err)
	}

	return &acc, nil
}

func (r *Repository) FindByAccountNumber(ctx context.Context, accountNumber string) (*account.Account, error) {
	query := `
		SELECT id, account_number, customer_id, balance, currency, status
		FROM accounts
		WHERE account_number = $1
	`

	var acc account.Account
	err := r.pool.QueryRow(ctx, query, accountNumber).Scan(
		&acc.ID,
		&acc.AccountNumber,
		&acc.CustomerID,
		&acc.Balance,
		&acc.Currency,
		&acc.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find account by account number: %w", err)
	}

	return &acc, nil
}

func (r *Repository) FindByCustomerID(ctx context.Context, customerID string) ([]*account.Account, error) {
	query := `
		SELECT id, account_number, customer_id, balance, currency, status
		FROM accounts
		WHERE customer_id = $1
	`

	rows, err := r.pool.Query(ctx, query, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to find accounts by customer ID: %w", err)
	}
	defer rows.Close()

	var accounts []*account.Account
	for rows.Next() {
		var acc account.Account
		err := rows.Scan(
			&acc.ID,
			&acc.AccountNumber,
			&acc.CustomerID,
			&acc.Balance,
			&acc.Currency,
			&acc.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan account: %w", err)
		}
		accounts = append(accounts, &acc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating accounts: %w", err)
	}

	return accounts, nil
}

func (r *Repository) Update(ctx context.Context, acc *account.Account) error {
	query := `
		UPDATE accounts
		SET balance = $1, currency = $2, status = $3
		WHERE id = $4
	`

	_, err := r.pool.Exec(ctx, query,
		acc.Balance,
		acc.Currency,
		acc.Status,
		acc.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	// Save events
	for _, event := range acc.GetEvents() {
		eventData, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %w", err)
		}

		_, err = r.pool.Exec(ctx, `
			INSERT INTO account_events (account_id, event_type, event_data)
			VALUES ($1, $2, $3)
		`, acc.ID, event.GetType(), eventData)
		if err != nil {
			return fmt.Errorf("failed to save event: %w", err)
		}
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM accounts WHERE id = $1`

	_, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	return nil
}
