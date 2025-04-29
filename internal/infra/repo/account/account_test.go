//go:build integration

package account

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

// setupTestDB creates a new PostgreSQL container, copy init scripts and returns a connection pool
func setupTestDB(t *testing.T, keepContainer bool) (*pgxpool.Pool, string) {
	ctx := context.Background()

	// Get the absolute path to the schema directory
	schemaDir, err := filepath.Abs("../../db/schema")
	require.NoError(t, err)

	dataDir, err := filepath.Abs("../../db/testdata")
	require.NoError(t, err)

	// Create PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "test",
		},

		WaitingFor: wait.ForAll(
			wait.ForListeningPort("5432/tcp"),
			wait.ForLog("database system is ready to accept connections"),
		),
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(schemaDir, "0001_customers_table.sql"),
				ContainerFilePath: "/docker-entrypoint-initdb.d/0001_customers.sql",
				FileMode:          0644,
			},
			{
				HostFilePath:      filepath.Join(dataDir, "0000_data.sql"),
				ContainerFilePath: "/docker-entrypoint-initdb.d/0003_customers_in_data.sql",
				FileMode:          0644,
			},
			{
				HostFilePath:      filepath.Join(schemaDir, "0002_accounts_table.sql"),
				ContainerFilePath: "/docker-entrypoint-initdb.d/0002_accounts.sql",
				FileMode:          0644,
			},
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	if !keepContainer {
		t.Cleanup(func() {
			require.NoError(t, container.Terminate(ctx))
		})
	} else {
		t.Logf("Container ID: %s", container.GetContainerID())
		t.Logf("Container will be kept running after test completion")
	}

	// Get container host and port
	host, err := container.Host(ctx)
	require.NoError(t, err)

	// Get the mapped port
	port, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	// Create connection string
	connString := "postgres://test:test@" + host + ":" + port.Port() + "/test?sslmode=disable"

	// Create connection pool
	config, err := pgxpool.ParseConfig(connString)
	require.NoError(t, err)
	config.MaxConns = 5
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool, connString
}

func TestAccountRepository_CreateAccount(t *testing.T) {
	// Set keepContainer to true to keep the container running after the test
	keepContainer := false
	pool, address := setupTestDB(t, keepContainer)

	log.Printf("container address: %s", address)

	repo := NewAccountRepository(pool)
	ctx := context.Background()

	tests := []struct {
		name    string
		account *account.Account
		wantErr bool
	}{
		{
			name: "successful account creation",
			account: account.NewAccount(
				"00000000-0000-0000-0000-000000000002",
				"00000000-0000-0000-0000-000000000000",
				"000000002",
				100.0,
				"USD",
			),
			wantErr: false,
		},
		{
			name: "duplicate account number",
			account: account.NewAccount(
				"00000000-0000-0000-0000-000000000001", // Same account number as above
				"00000000-0000-0000-0000-000000000000",
				"00000001",
				200.0,
				"USD",
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := repo.CreateAccount(ctx, tt.account)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, acc)
			}
		})
	}
}
