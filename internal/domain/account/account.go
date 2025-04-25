package account

import (
	"errors"
	"time"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// AccountType represents the type of bank account
type AccountType string

const (
	AccountTypeSavings  AccountType = "savings"
	AccountTypeChecking AccountType = "checking"
	AccountTypeLoan     AccountType = "loan"
)

// Account represents a bank account in the system.
// It is an aggregate root that manages its own balance and state.
type Account struct {
	// AccountBase ??
	ID            string        // Unique identifier of the account, must be in UUID format
	AccountNumber string        // Account number (e.g., 1234567890)
	Balance       float64       // Current balance of the account
	Status        AccountStatus // Current status of the account (active/blocked)
	Currency      string        // Currency code (e.g., USD, EUR)
	CreatedAt     time.Time     // When the account was created
	UpdatedAt     time.Time     // When the account was last updated
	events        []Event       // List of domain events that occurred on this account
}

// AccountStatus represents the possible states of an account
type AccountStatus string

const (
	AccountStatusInactive AccountStatus = "inactive" // Account is active and can perform transactions
	AccountStatusActive   AccountStatus = "active"   // Account is active and can perform transactions
	AccountStatusBlocked  AccountStatus = "blocked"  // Account is blocked and cannot perform transactions
)

// NewAccount creates a new account with the given ID and initial balance.
// It automatically sets the account status to active and records the creation event.
func NewAccount(id, number string, initialBalance float64, currency string) *Account {
	now := time.Now().UTC()
	account := &Account{
		ID:            id,
		AccountNumber: number,
		Balance:       initialBalance,
		Currency:      currency,
		Status:        AccountStatusActive,
		CreatedAt:     now,
		UpdatedAt:     now,
		events:        make([]Event, 0),
	}

	account.addEvent(AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          generateEventID(),
			Type:        AccountCreatedEventType,
			AggregateID: id,
			CreatedAt:   now,
		},
		InitialBalance: initialBalance,
	})

	return account
}

// Block marks the account as blocked, preventing any transactions.
// It updates the account status and records a blocking event.
func (a *Account) Block() {
	now := time.Now().UTC()
	a.UpdatedAt = now
	a.Status = AccountStatusBlocked

	a.addEvent(AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          generateEventID(),
			Type:        AccountBlockedEventType,
			AggregateID: a.ID,
			CreatedAt:   now,
		},
	})
}

// Unblock marks the account as active, allowing transactions again.
// It updates the account status and records an unblocking event.
func (a *Account) Unblock() {
	now := time.Now().UTC()
	a.UpdatedAt = now
	a.Status = AccountStatusActive

	a.addEvent(AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          generateEventID(),
			Type:        AccountUnblockedEventType,
			AggregateID: a.ID,
			CreatedAt:   now,
		},
	})
}

// Deposit adds the specified amount to the account balance.
// It updates the account's balance and records a deposit event.
func (a *Account) Deposit(amount float64) {
	now := time.Now().UTC()
	a.Balance += amount
	a.UpdatedAt = now

	a.addEvent(FundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          generateEventID(),
			Type:        AccountFundsDepositedEventType,
			AggregateID: a.ID,
			CreatedAt:   now,
		},
		Amount:   amount,
		Balance:  a.Balance,
		Currency: a.Currency,
	})
}

// Withdraw subtracts the specified amount from the account balance.
// It returns an error if there are insufficient funds.
// On success, it updates the balance and records a withdrawal event.
func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	now := time.Now().UTC()
	a.Balance -= amount
	a.UpdatedAt = now

	a.addEvent(FundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          generateEventID(),
			Type:        AccountFundsWithdrawnEventType,
			AggregateID: a.ID,
			CreatedAt:   now,
		},
		Amount:   amount,
		Balance:  a.Balance,
		Currency: a.Currency,
	})

	return nil
}

// GetEvents returns all domain events that have occurred on this account.
func (a *Account) GetEvents() []Event {
	return a.events
}

// ClearEvents removes all recorded events from the account.
// This is typically called after events have been processed.
func (a *Account) ClearEvents() {
	a.events = make([]Event, 0)
}

// addEvent is an internal method to record a new domain event.
func (a *Account) addEvent(event Event) {
	a.events = append(a.events, event)
}
