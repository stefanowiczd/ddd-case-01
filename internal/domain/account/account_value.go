package account

// AccountStatus represents the possible states of an account
type AccountStatus string

const (
	AccountStatusInactive AccountStatus = "inactive" // Account is active and can perform transactions
	AccountStatusActive   AccountStatus = "active"   // Account is active and can perform transactions
	AccountStatusBlocked  AccountStatus = "blocked"  // Account is blocked and cannot perform transactions
)

func (s AccountStatus) String() string {
	return string(s)
}

// IsValid checks if the account status is valid
func (s AccountStatus) IsValid() bool {
	return s == AccountStatusInactive || s == AccountStatusActive || s == AccountStatusBlocked
}
