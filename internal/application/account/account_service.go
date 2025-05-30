package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	accountdomain "github.com/stefanowiczd/ddd-case-01/internal/domain/account"
)

type Account = accountdomain.Account

// AccountService handles account-related use cases
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

// CreateAccountDTO represents the data needed to create a new account
type CreateAccountDTO struct {
	CustomerID     string  `json:"customerId"`
	InitialBalance float64 `json:"initialBalance"`
	Currency       string  `json:"currency"`
}

// AccountResponseDTO represents the account data returned to clients
type AccountResponseDTO struct {
	ID            string  `json:"id"`
	AccountNumber string  `json:"accountNumber"`
	CustomerID    string  `json:"customerId"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}

type CreateAccountResponseDTO struct {
	AccountResponseDTO
}

// CreateAccount creates a new account
func (s *AccountService) CreateAccount(ctx context.Context, dto CreateAccountDTO) (CreateAccountResponseDTO, error) {
	if dto.InitialBalance < 0 {
		return CreateAccountResponseDTO{}, ErrInvalidInitialBalanceAmount
	}

	accountNumber := generateAccountNumber()
	id := uuid.New()

	if _, err := s.accountQueryRepo.FindByID(ctx, id); err != nil {
		if !errors.Is(err, accountdomain.ErrAccountNotFound) {
			return CreateAccountResponseDTO{}, fmt.Errorf("finding account by id: %w", err)
		}
	}

	account := accountdomain.NewAccount(id, uuid.MustParse(dto.CustomerID), accountNumber, dto.InitialBalance, dto.Currency)

	if err := s.accountEventRepo.CreateEvents(ctx, account.GetEvents()); err != nil {
		return CreateAccountResponseDTO{}, err
	}

	return CreateAccountResponseDTO{
		AccountResponseDTO: ToDTO(account),
	}, nil
}

type GetAccountDTO struct {
	AccountID uuid.UUID `json:"accountId"`
}

// GetAccount retrieves an account by its ID
func (s *AccountService) GetAccount(ctx context.Context, dto GetAccountDTO) (AccountResponseDTO, error) {
	account, err := s.accountQueryRepo.FindByID(ctx, dto.AccountID)
	if err != nil {
		if errors.Is(err, accountdomain.ErrAccountNotFound) {
			return AccountResponseDTO{}, fmt.Errorf("finding account by id: %w", ErrAccountNotFound)
		}
		return AccountResponseDTO{}, fmt.Errorf("finding account by id: %w", err)
	}

	return ToDTO(account), nil
}

// DepositDTO represents the data needed to deposit money
type DepositDTO struct {
	AccountID uuid.UUID `json:"accountId"`
	Amount    float64   `json:"amount"`
}

// Deposit adds money to an account
func (s *AccountService) Deposit(ctx context.Context, dto DepositDTO) error {
	if dto.Amount <= 0 {
		return ErrInvalidDepositAmount
	}

	account, err := s.accountQueryRepo.FindByID(ctx, dto.AccountID)
	if err != nil {
		if errors.Is(err, accountdomain.ErrAccountNotFound) {
			return fmt.Errorf("finding account by id: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("finding account by id: %w", err)
	}

	account.Deposit(dto.Amount)

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// WithdrawDTO represents the data needed to withdraw money
type WithdrawDTO struct {
	AccountID uuid.UUID `json:"accountId"`
	Amount    float64   `json:"amount"`
}

// Withdraw removes money from an account
func (s *AccountService) Withdraw(ctx context.Context, dto WithdrawDTO) error {
	if dto.Amount <= 0 {
		return ErrInvalidWithdrawAmount
	}

	account, err := s.accountQueryRepo.FindByID(ctx, dto.AccountID)
	if err != nil {
		if errors.Is(err, accountdomain.ErrAccountNotFound) {
			return fmt.Errorf("finding account by id: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("finding account by id: %w", err)
	}

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// BlockAccountDTO represents the data needed to block an account
type BlockAccountDTO struct {
	AccountID uuid.UUID `json:"accountId"`
}

// BlockAccount blocks an account
func (s *AccountService) BlockAccount(ctx context.Context, dto BlockAccountDTO) error {
	account, err := s.accountQueryRepo.FindByID(ctx, dto.AccountID)
	if err != nil {
		if errors.Is(err, accountdomain.ErrAccountNotFound) {
			return fmt.Errorf("finding account by id: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("finding account by id: %w", err)
	}

	account.Block()

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// UnblockAccountDTO represents the data needed to unblock an account
type UnblockAccountDTO struct {
	AccountID uuid.UUID `json:"accountId"`
}

// UnblockAccount unblocks an account
func (s *AccountService) UnblockAccount(ctx context.Context, dto UnblockAccountDTO) error {
	account, err := s.accountQueryRepo.FindByID(ctx, dto.AccountID)
	if err != nil {
		if errors.Is(err, accountdomain.ErrAccountNotFound) {
			return fmt.Errorf("finding account by id: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("finding account by id: %w", err)
	}

	account.Unblock()

	return s.accountEventRepo.CreateEvents(ctx, account.GetEvents())
}

// ToDTO converts an Account domain model to AccountResponseDTO
func ToDTO(account *Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:            account.ID.String(),
		AccountNumber: account.AccountNumber,
		CustomerID:    account.CustomerID.String(),
		Balance:       account.Balance,
		Currency:      account.Currency,
		Status:        account.Status.String(),
		CreatedAt:     account.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:     account.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// ToDTOList converts a list of Account domain models to AccountResponseDTOs
func ToDTOList(accounts []*Account) []AccountResponseDTO {
	dtos := make([]AccountResponseDTO, len(accounts))
	for i, account := range accounts {
		dtos[i] = ToDTO(account)
	}
	return dtos
}

type GetCustomerAccountsDTO struct {
	CustomerID uuid.UUID `json:"customerId"`
}

type GetCustomerAccountsResponseDTO struct {
	Accounts []AccountResponseDTO `json:"accounts"`
}

// TODO: move to customer service where responsibility of customer is
// GetCustomerAccounts retrieves all accounts for a customer
func (s *AccountService) GetCustomerAccounts(ctx context.Context, dto GetCustomerAccountsDTO) (GetCustomerAccountsResponseDTO, error) {
	_, err := s.customerQueryRepo.FindByID(ctx, dto.CustomerID)
	if err != nil {
		return GetCustomerAccountsResponseDTO{}, ErrCustomerNotFound
	}

	accounts, err := s.accountQueryRepo.FindByCustomerID(ctx, dto.CustomerID)
	if err != nil {
		return GetCustomerAccountsResponseDTO{}, err
	}

	return GetCustomerAccountsResponseDTO{
		Accounts: ToDTOList(accounts),
	}, nil
}

// generateAccountNumber generates a unique account number
func generateAccountNumber() string {
	return uuid.New().String()
}
