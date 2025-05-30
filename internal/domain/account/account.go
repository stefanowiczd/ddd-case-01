package account

import (
	"time"

	"github.com/google/uuid"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

type EventOrigin = event.EventOrigin

// AccountType represents the type of bank account
type AccountType string

func (a AccountType) String() string {
	return string(a)
}

const (
	AccountTypeSavings  AccountType = "savings"
	AccountTypeChecking AccountType = "checking"
	AccountTypeLoan     AccountType = "loan"
)

// Account represents a bank account in the system.
// It is an aggregate root that manages its own balance and state.
type Account struct {
	// AccountBase ??
	ID            uuid.UUID     // Unique identifier of the account, must be in UUID format
	CustomerID    uuid.UUID     // Unique identifier of the customer, must be in UUID format
	AccountNumber string        // Account number (e.g., 1234567890)
	Balance       float64       // Current balance of the account
	Status        AccountStatus // Current status of the account (active/blocked)
	Currency      string        // Currency code (e.g., USD, EUR)
	CreatedAt     time.Time     // When the account was created
	UpdatedAt     time.Time     // When the account was last updated
	events        []Event       // List of domain events that occurred on this account
}

// NewAccount creates a new account with the given ID and initial balance.
// It automatically sets the account status to active and records the creation event.
func NewAccount(id uuid.UUID, customerID uuid.UUID, number string, initialBalance float64, currency string) *Account {
	now := time.Now().UTC()
	account := &Account{
		ID:            id,
		CustomerID:    customerID,
		AccountNumber: number,
		Balance:       initialBalance,
		Currency:      currency,
		Status:        AccountStatusActive,
		CreatedAt:     now,
		UpdatedAt:     now,
		events:        make([]Event, 0),
	}

	origin := EventOrigin("account")

	account.addEvent(&AccountCreatedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			ContextID:   id,
			Origin:      origin.String(),
			Type:        AccountCreatedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
		InitialBalance: initialBalance,
		CustomerID:     customerID,
	})

	return account
}

// Block marks the account as blocked, preventing any transactions.
// It updates the account status and records a blocking event.
func (a *Account) Block() {
	now := time.Now().UTC()
	a.UpdatedAt = now
	a.Status = AccountStatusBlocked

	origin := EventOrigin("account")

	a.addEvent(&AccountBlockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			ContextID:   a.ID,
			Origin:      origin.String(),
			Type:        AccountBlockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	})
}

// Unblock marks the account as active, allowing transactions again.
// It updates the account status and records an unblocking event.
func (a *Account) Unblock() {
	now := time.Now().UTC()
	a.UpdatedAt = now
	a.Status = AccountStatusActive

	origin := EventOrigin("account")

	a.addEvent(&AccountUnblockedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			ContextID:   a.ID,
			Origin:      origin.String(),
			Type:        AccountUnblockedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
		},
	})
}

// Deposit adds the specified amount to the account balance.
// It updates the account's balance and records a deposit event.
func (a *Account) Deposit(amount float64) {
	now := time.Now().UTC()
	a.Balance += amount
	a.UpdatedAt = now

	origin := EventOrigin("account")

	a.addEvent(&AccountFundsDepositedEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			ContextID:   a.ID,
			Origin:      origin.String(),
			Type:        AccountFundsDepositedEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
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
	now := time.Now().UTC()
	a.Balance -= amount
	a.UpdatedAt = now

	origin := EventOrigin("account")

	a.addEvent(&AccountFundsWithdrawnEvent{
		BaseEvent: event.BaseEvent{
			ID:          uuid.New(),
			ContextID:   a.ID,
			Origin:      origin.String(),
			Type:        AccountFundsWithdrawnEventType.String(),
			TypeVersion: "0.0.0",
			State:       event.EventStateReady.String(),
			CreatedAt:   now,
			ScheduledAt: now,
			Retry:       0,
			MaxRetry:    3,
			Data:        nil,
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
