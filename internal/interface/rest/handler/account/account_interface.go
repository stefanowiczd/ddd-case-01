package account

import (
	"context"

	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
)

//go:generate mockgen -destination=./mock/account_handler_mock.go -package=mock -source=./account_interface.go

// AccountQueryService defines the contract for account query operations
type AccountQueryService interface {
	// GetAccount retrieves an account by its ID
	GetAccount(ctx context.Context, dto applicationaccount.GetAccountDTO) (applicationaccount.AccountResponseDTO, error)

	// GetCustomerAccounts retrieves all accounts for a customer
	GetCustomerAccounts(ctx context.Context, dto applicationaccount.GetCustomerAccountsDTO) (applicationaccount.GetCustomerAccountsResponseDTO, error)
}

// AccountService defines the contract for account operations
type AccountService interface {
	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, dto applicationaccount.CreateAccountDTO) (applicationaccount.CreateAccountResponseDTO, error)

	// Deposit adds money to an account
	Deposit(ctx context.Context, dto applicationaccount.DepositDTO) error

	// Withdraw removes money from an account
	Withdraw(ctx context.Context, dto applicationaccount.WithdrawDTO) error

	// BlockAccount blocks an account
	BlockAccount(ctx context.Context, dto applicationaccount.BlockAccountDTO) error

	// UnblockAccount unblocks an account
	UnblockAccount(ctx context.Context, dto applicationaccount.UnblockAccountDTO) error
}
