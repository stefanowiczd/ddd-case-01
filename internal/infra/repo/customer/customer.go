package customer

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/infra/repo/query"
)

// CustomerRepository is a repository for customer operations
type CustomerRepository struct {
	Conn *pgxpool.Pool
	Q    *query.Queries
}

func NewCustomerRepository(conn *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{
		Conn: conn,
		Q:    query.New(conn),
	}
}

// FindByID finds a customer by id
func (r *CustomerRepository) FindByID(ctx context.Context, id uuid.UUID) (*customerdomain.Customer, error) {
	customer, err := r.Q.FindCustomerByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("finding customer by id: %w", customerdomain.ErrCustomerNotFound)
		}

		return nil, fmt.Errorf("finding customer by id: %w", err)
	}

	return &customerdomain.Customer{
		ID:          customer.ID.Bytes,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		Phone:       customer.Phone.String,
		DateOfBirth: customer.DateOfBirth,
		Address: customerdomain.Address{
			Street:     customer.AddressStreet.String,
			City:       customer.AddressCity.String,
			State:      customer.AddressState.String,
			PostalCode: customer.AddressZipCode.String,
			Country:    customer.AddressCountry.String,
		},
		Status:    customerdomain.CustomerStatus(customer.Status),
		Accounts:  []string{},
		CreatedAt: customer.CreatedAt.Time,
		UpdatedAt: customer.UpdatedAt.Time,
	}, nil
}

// CreateCustomer creates a new customer
func (r *CustomerRepository) CreateCustomer(ctx context.Context, cust *customerdomain.Customer) (*customerdomain.Customer, error) {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting transaction: create customer: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	customer, err := r.Q.CreateCustomer(
		ctx,
		query.CreateCustomerParams{
			ID:             pgtype.UUID{Bytes: cust.ID, Valid: true},
			FirstName:      cust.FirstName,
			LastName:       cust.LastName,
			Email:          cust.Email,
			Phone:          pgtype.Text{String: cust.Phone, Valid: true},
			DateOfBirth:    cust.DateOfBirth,
			AddressStreet:  pgtype.Text{String: cust.Address.Street, Valid: true},
			AddressCity:    pgtype.Text{String: cust.Address.City, Valid: true},
			AddressState:   pgtype.Text{String: cust.Address.State, Valid: true},
			AddressZipCode: pgtype.Text{String: cust.Address.PostalCode, Valid: true},
			AddressCountry: pgtype.Text{String: cust.Address.Country, Valid: true},
			Status:         cust.Status.String(),
			CreatedAt:      pgtype.Timestamp{Time: cust.CreatedAt, Valid: true},
			UpdatedAt:      pgtype.Timestamp{Time: cust.UpdatedAt, Valid: true},
		})
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr):
			switch pgErr.Code {
			case query.POSTGRESQL_DUPLICATE_KEY_CODE:
				return nil, fmt.Errorf("checking db error: executing query: create customer: %w", customerdomain.ErrCustomerAlreadyExists)
			default:
				return nil, fmt.Errorf("checking db error: executing query: create customer: %w", err)
			}
		}
		return nil, fmt.Errorf("executing query: create customer: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("committing transaction: create customer: %w", err)
	}

	return &customerdomain.Customer{
		ID:          customer.ID.Bytes,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		Phone:       customer.Phone.String,
		DateOfBirth: customer.DateOfBirth,
		Address: customerdomain.Address{
			Street:     customer.AddressStreet.String,
			City:       customer.AddressCity.String,
			State:      customer.AddressState.String,
			PostalCode: customer.AddressZipCode.String,
			Country:    customer.AddressCountry.String,
		},
		Status:    customerdomain.CustomerStatus(customer.Status),
		Accounts:  []string{},
		CreatedAt: customer.CreatedAt.Time,
		UpdatedAt: customer.UpdatedAt.Time,
	}, nil
}
