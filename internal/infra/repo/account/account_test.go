//go:build integration

package account

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

func TestAccountRepository_FindByID_NoRows(t *testing.T) {
	ctx := context.Background()
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountRepository(pool)

	acc, err := repo.FindByID(ctx, uuid.New())

	require.ErrorIs(t, err, account.ErrAccountNotFound)
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
