//go:build integration

package account

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

func TestAccountRepository_FindByID_NoRows(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountRepository(pool)

	acc, err := repo.FindByID(ctx, uuid.New())

	require.ErrorIs(t, err, accountdomain.ErrAccountNotFound)
	require.Nil(t, acc)

	accs, err := repo.FindByCustomerID(ctx, uuid.New())
	require.Len(t, accs, 0)
}

func TestAccountRepository_FindBy(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountRepository(pool)

	eventID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	acc, err := repo.FindByID(ctx, eventID)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, acc)

	customerID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	accs, err := repo.FindByCustomerID(ctx, customerID)
	require.Len(t, accs, 1)
}

func TestAccountRepository_CreateAccount(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountRepository(pool)

	a := &accountdomain.Account{
		ID:            uuid.MustParse("00000000-0000-0000-0000-111111111111"),
		CustomerID:    uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		AccountNumber: "1111111111",
		Balance:       1000,
		Currency:      "USD",
		Status:        accountdomain.AccountStatusActive,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	acc, err := repo.CreateAccount(ctx, a)

	require.NoError(t, err)
	require.NotNil(t, acc)

	acc, err = repo.FindByID(ctx, uuid.MustParse("00000000-0000-0000-0000-111111111111"))
	require.NoError(t, err)
	require.NotNil(t, acc)

	_, err = repo.CreateAccount(ctx, a)
	require.ErrorIs(t, err, accountdomain.ErrAccountAlreadyExists)
}
