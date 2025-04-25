//go:build unit

package account

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func testAccountNumber() string {
	return uuid.New().String()
}

func testCustomerID() string {
	return uuid.New().String()
}

func Test_NewAccount(t *testing.T) {
	id := testCustomerID()

	type testCaseParams struct {
		accountID     string
		accountNumber string
	}

	type testCaseExpected struct {
		accountID        string
		accountStatus    AccountStatus
		accountBalance   float64
		accountCurrency  string
		eventsNumber     int
		eventType        string
		eventAggregateID string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "should create new account with active status",
			params: testCaseParams{
				accountID:     id,
				accountNumber: testAccountNumber(),
			},
			expected: testCaseExpected{
				accountID:        id,
				accountBalance:   0.0,
				accountCurrency:  "USD",
				accountStatus:    AccountStatusActive,
				eventsNumber:     1,
				eventType:        AccountCreatedEventType,
				eventAggregateID: id,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(tt.params.accountID, tt.params.accountNumber, 0, "USD")

			// Account checks
			require.Equal(t, tt.expected.accountID, account.ID)
			require.Equal(t, tt.expected.accountStatus, account.Status)
			require.Equal(t, tt.expected.accountBalance, account.Balance)
			require.Equal(t, tt.expected.accountCurrency, account.Currency)

			// Event checks
			require.Len(t, account.events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, account.events[0].GetType())
			require.Equal(t, tt.expected.eventAggregateID, account.events[0].GetAggregateID())
			require.Less(t, account.events[0].GetCreatedAt(), time.Now())

		})
	}
}

func Test_Account_Block(t *testing.T) {

	type testCaseParams struct{}

	type testCaseExpected struct {
		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name:   "should create a new event when blocking an account",
			params: testCaseParams{},
			expected: testCaseExpected{
				eventsNumber: 2,
				eventType:    AccountBlockedEventType,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(testCustomerID(), testAccountNumber(), 0, "USD")

			account.Block()
			require.Len(t, account.events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, account.events[1].GetType())
		})
	}
}

func Test_Account_Unblock(t *testing.T) {

	type testCaseParams struct{}

	type testCaseExpected struct {
		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name:   "should create a new event when unblocking an account",
			params: testCaseParams{},
			expected: testCaseExpected{
				eventsNumber: 2,
				eventType:    AccountUnblockedEventType,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(testCustomerID(), testAccountNumber(), 0, "USD")

			account.Unblock()
			require.Len(t, account.events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, account.events[1].GetType())
		})
	}
}

func Test_Account_Deposit(t *testing.T) {

	type testCaseParams struct{}

	type testCaseExpected struct {
		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name:   "should create a new event when depositing funds",
			params: testCaseParams{},
			expected: testCaseExpected{
				eventsNumber: 2,
				eventType:    AccountFundsDepositedEventType,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(testCustomerID(), testAccountNumber(), 0, "USD")

			account.Deposit(404)
			require.Len(t, account.events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, account.events[1].GetType())
		})
	}
}

func Test_Account_Withdraw(t *testing.T) {

	type testCaseParams struct{}

	type testCaseExpected struct {
		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name:   "should create a new event when withdrawing funds",
			params: testCaseParams{},
			expected: testCaseExpected{
				eventsNumber: 2,
				eventType:    AccountFundsWithdrawnEventType,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := NewAccount(testCustomerID(), testAccountNumber(), 1000, "USD")

			account.Withdraw(404)
			require.Len(t, account.events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, account.events[1].GetType())
		})
	}
}
