package account

import (
	"github.com/google/uuid"

	"github.com/stefanowiczd/ddd-case-01/internal/domain/event"
)

// AccountCreatedEvent is emitted when a new account is created
type AccountCreatedEvent struct {
	event.BaseEvent
	CustomerID     uuid.UUID `json:"customer_id"`
	InitialBalance float64   `json:"initial_balance"` // The initial balance of the account
}

// AccountFundsDepositedEvent is emitted when funds are deposited into an account
type AccountFundsDepositedEvent struct {
	event.BaseEvent
	Amount   float64 `json:"amount"`   // The amount that was deposited
	Balance  float64 `json:"balance" ` // The new balance after the deposit
	Currency string  `json:"currency"` // The currency of the account
}

// AccountFundsWithdrawnEvent is emitted when funds are withdrawn from an account
type AccountFundsWithdrawnEvent struct {
	event.BaseEvent
	Amount   float64 `json:"amount"`   // The amount that was withdrawn
	Balance  float64 `json:"balance"`  // The new balance after the withdrawal
	Currency string  `json:"currency"` // The currency of the account
}

// AccountBlockedEvent is emitted when an account is blocked
type AccountBlockedEvent struct {
	event.BaseEvent
}

// AccountUnblockedEvent is emitted when an account is unblocked
type AccountUnblockedEvent struct {
	event.BaseEvent
}
