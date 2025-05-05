package account

// AccountEventType represents the type of account event
type AccountEventType string

// String returns the string representation of the account event type
func (e AccountEventType) String() string {
	return string(e)
}

const (
	AccountCreatedEventType   AccountEventType = "account.created"
	AccountActivatedEventType AccountEventType = "account.activated"

	AccountBlockedEventType   AccountEventType = "account.blocked"
	AccountUnblockedEventType AccountEventType = "account.unblocked"

	AccountFundsDepositedEventType AccountEventType = "account.funds.deposited"
	AccountFundsWithdrawnEventType AccountEventType = "account.funds.withdrawn"
)
