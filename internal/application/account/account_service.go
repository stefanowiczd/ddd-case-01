package account

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/account"
	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrCustomerNotFound = errors.New("customer not found")

	ErrInvalidAmount = errors.New("invalid amount")
)

type Account = accountdomain.Account

// Service handles account-related use cases
type AccountService struct {
	accountQueryRepo  AccountQueryRepository
	customerQueryRepo CustomerQueryRepository

	accountEventRepo AccountEventRepository
}

// NewService creates a new account service
func NewService(
	accountQueryRepo AccountQueryRepository,
	customerQueryRepo CustomerQueryRepository,
	accountEventRepo AccountEventRepository) *AccountService {
	return &AccountService{
		accountQueryRepo:  accountQueryRepo,
		accountEventRepo:  accountEventRepo,
		customerQueryRepo: customerQueryRepo,
	}
}

// CreateAccount creates a new account
func (s *AccountService) CreateAccount(ctx context.Context, customerID string, initialBalance float64, currency string) (*Account, error) {
	if initialBalance < 0 {
		return nil, ErrInvalidAmount
	}

	accountNumber := generateAccountNumber()
	account := account.NewAccount(accountNumber, customerID, initialBalance, currency)

	if err := s.accountEventRepo.CreateEvents(ctx, account.GetEvents()); err != nil {
		return nil, err
	}

	return account, nil
}

// GetAccount retrieves an account by its ID
func (s *AccountService) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	account, err := s.accountQueryRepo.FindByID(ctx, accountID)
	if err != nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

// Deposit adds money to an account
func (s *AccountService) Deposit(ctx context.Context, accountID string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	account, err := s.accountQueryRepo.FindByID(ctx, accountID)
	if err != nil {
		return ErrAccountNotFound
	}

	account.Deposit(amount)

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// Withdraw removes money from an account
func (s *AccountService) Withdraw(ctx context.Context, accountID string, amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	account, err := s.accountQueryRepo.FindByID(ctx, accountID)
	if err != nil {
		return ErrAccountNotFound
	}

	// TODO: add validation for the account balance amount???

	account.Withdraw(amount)

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// BlockAccount blocks an account
func (s *AccountService) BlockAccount(ctx context.Context, accountID string) error {
	account, err := s.accountQueryRepo.FindByID(ctx, accountID)
	if err != nil {
		return ErrAccountNotFound
	}

	account.Block()

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// UnblockAccount unblocks an account
func (s *AccountService) UnblockAccount(ctx context.Context, accountID string) error {
	account, err := s.accountQueryRepo.FindByID(ctx, accountID)
	if err != nil {
		return ErrAccountNotFound
	}

	account.Unblock()

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// GetCustomerAccounts retrieves all accounts for a customer
func (s *AccountService) GetCustomerAccounts(ctx context.Context, customerID string) ([]*account.Account, error) {
	_, err := s.customerQueryRepo.FindByID(ctx, customerID)
	if err != nil {
		return nil, ErrCustomerNotFound
	}

	return s.accountQueryRepo.FindByCustomerID(ctx, customerID)
}

// generateAccountNumber generates a unique account number
func generateAccountNumber() string {
	return uuid.New().String()
}
