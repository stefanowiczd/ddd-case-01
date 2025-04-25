package account

import (
	"time"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

const (
	AccountCreatedEventType   = "account.created"
	AccountActivatedEventType = "account.activated"
	FundsDepositedEventType   = "account.funds.deposited"
	FundsWithdrawnEventType   = "account.funds.withdrawn"
	AccountBlockedEventType   = "account.blocked"
	AccountUnblockedEventType = "account.unblocked"
)

// AccountCreatedEvent is emitted when a new account is created
type AccountCreatedEvent struct {
	event.BaseEvent
	InitialBalance float64 // The initial balance of the account
}

// FundsDepositedEvent is emitted when funds are deposited into an account
type FundsDepositedEvent struct {
	event.BaseEvent
	Amount   float64 // The amount that was deposited
	Balance  float64 // The new balance after the deposit
	Currency string  // The currency of the account
}

// FundsWithdrawnEvent is emitted when funds are withdrawn from an account
type FundsWithdrawnEvent struct {
	event.BaseEvent
	Amount   float64 // The amount that was withdrawn
	Balance  float64 // The new balance after the withdrawal
	Currency string  // The currency of the account
}

// AccountBlockedEvent is emitted when an account is blocked
type AccountBlockedEvent struct {
	event.BaseEvent
}

// AccountUnblockedEvent is emitted when an account is unblocked
type AccountUnblockedEvent struct {
	event.BaseEvent
}

// generateEventID creates a unique identifier for an event based on the current timestamp
func generateEventID() string {
	return time.Now().Format("20060102150405.000000000")
}
