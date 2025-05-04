//go:build integration

package customer

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	customerdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/customer"
)

func TestCustomerRepository_FindByID_NoRows(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewCustomerRepository(pool)

	cust, err := repo.FindByID(ctx, uuid.New())

	require.ErrorIs(t, err, customerdomain.ErrCustomerNotFound)
	require.Nil(t, cust)
}

func TestCustomerRepository_CreateCustomer(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewCustomerRepository(pool)

	cust := customerdomain.Customer{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe.2@example.com",
		Phone:     "1234567890",
		Address: customerdomain.Address{
			Street:     "Street 1",
			City:       "Warsaw",
			State:      "Masovian",
			PostalCode: "00-000",
			Country:    "Poland",
		},
		Status:    customerdomain.CustomerStatusActive,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	custOut, err := repo.CreateCustomer(ctx, &cust)
	require.NoError(t, err)
	require.NotNil(t, custOut)

	require.Equal(t, cust.ID, custOut.ID)
	require.Equal(t, cust.FirstName, custOut.FirstName)
	require.Equal(t, cust.LastName, custOut.LastName)
	require.Equal(t, cust.Email, custOut.Email)
	require.Equal(t, cust.Phone, custOut.Phone)
	require.Equal(t, cust.DateOfBirth, custOut.DateOfBirth)
	require.Equal(t, cust.Address, custOut.Address)
	require.Equal(t, cust.Status, custOut.Status)
	require.Equal(t, cust.CreatedAt, custOut.CreatedAt)
	require.Equal(t, cust.UpdatedAt, custOut.UpdatedAt)

	_, err = repo.CreateCustomer(ctx, &cust)
	require.ErrorIs(t, err, customerdomain.ErrCustomerAlreadyExists)
}
