//go:build unit

package customer

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_NewCustomer(t *testing.T) {

	type testCaseParams struct {
		customerID uuid.UUID
		firstName  string
		lastName   string
		phone      string
		email      string
		address    Address
	}

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
			name: "should create new cutomer with active status",
			params: testCaseParams{
				customerID: uuid.New(),
				firstName:  "John",
				lastName:   "Doe",
				phone:      "1234567890",
				email:      "john.doe@example.com",
				address:    Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
			},
			expected: testCaseExpected{
				eventsNumber: 1,
				eventType:    CustomerCreatedEventType.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := NewCustomer(tt.params.customerID, tt.params.firstName, tt.params.lastName, tt.params.phone, tt.params.email, tt.params.address)

			// Event checks
			require.Len(t, customer.Events, tt.expected.eventsNumber)
			require.Equal(t, tt.expected.eventType, customer.Events[0].GetType())
			require.Equal(t, tt.params.customerID, customer.Events[0].GetContextID())
			require.Less(t, customer.Events[0].GetCreatedAt(), time.Now().UTC())

		})
	}
}

func Test_Customer_Update_Name(t *testing.T) {

	type testCaseParams struct {
		updateType CustomerEventType
		firstName  string
		lastName   string
		phone      string
		email      string
		address    Address
	}

	type testCaseExpected struct {
		updateType   CustomerEventType
		firstName    string
		lastName     string
		phone        string
		email        string
		address      Address
		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "should update customer name",
			params: testCaseParams{
				updateType: CustomerUpdatedAllEventType,
				firstName:  "John Second",
				lastName:   "Doe Second",
				phone:      "0987654321",
				email:      "jane.doe@example.com",
				address:    Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
			},
			expected: testCaseExpected{
				updateType: CustomerUpdatedAllEventType,
				firstName:  "John Second",
				lastName:   "Doe Second",
				phone:      "0987654321",
				email:      "jane.doe@example.com",
				address:    Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},

				eventsNumber: 2, // 1 for creation and 1 for update
				eventType:    CustomerUpdatedAllEventType.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := NewCustomer(
				uuid.New(),
				"John",
				"Doe",
				"1234567890",
				"john.doe@example.com",
				Address{Street: "Street 1111", City: "Warsaw 2", State: "Masovian 4", PostalCode: "11-111", Country: "USA"},
			)

			customer.Update(tt.params.updateType, tt.params.firstName, tt.params.lastName, tt.params.phone, tt.params.email, tt.params.address)

			require.Equal(t, tt.expected.firstName, customer.FirstName)
			require.Equal(t, tt.expected.lastName, customer.LastName)
			require.Equal(t, tt.expected.phone, customer.Phone)
			require.Equal(t, tt.expected.email, customer.Email)
			require.True(t, customer.Address.compare(tt.expected.address))

			require.Equal(t, tt.expected.eventsNumber, len(customer.Events))
			require.Equal(t, tt.expected.eventType, customer.Events[1].GetType())
		})
	}
}

func Test_Customer_Block(t *testing.T) {

	type testCaseParams struct {
		reason string
	}

	type testCaseExpected struct {
		customerStatus CustomerStatus

		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "should block customer",
			params: testCaseParams{
				reason: "customer blocked account",
			},
			expected: testCaseExpected{
				customerStatus: CustomerStatusBlocked,

				eventsNumber: 2,
				eventType:    CustomerBlockedEventType.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := NewCustomer(
				uuid.New(),
				"John",
				"Doe",
				"1234567890",
				"john.doe@example.com",
				Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
			)

			customer.Block(tt.params.reason)

			require.Equal(t, tt.expected.customerStatus, customer.Status)
			require.Equal(t, tt.expected.eventsNumber, len(customer.Events))
			require.Equal(t, tt.expected.eventType, customer.Events[1].GetType())
		})
	}
}

func Test_Customer_Unblock(t *testing.T) {

	type testCaseParams struct {
		reason string
	}

	type testCaseExpected struct {
		customerStatus CustomerStatus

		eventsNumber int
		eventType    string
	}

	tests := []struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}{
		{
			name: "should unblock customer",
			params: testCaseParams{
				reason: "customer unblocked account",
			},
			expected: testCaseExpected{
				customerStatus: CustomerStatusActive,

				eventsNumber: 2,
				eventType:    CustomerUnblockedEventType.String(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customer := NewCustomer(
				uuid.New(),
				"John",
				"Doe",
				"1234567890",
				"john.doe@example.com",
				Address{Street: "Street 1", City: "Warsaw", State: "Masovian", PostalCode: "00-000", Country: "Poland"},
			)

			customer.Unblock()

			require.Equal(t, tt.expected.customerStatus, customer.Status)

			require.Equal(t, tt.expected.eventsNumber, len(customer.Events))
			require.Equal(t, tt.expected.eventType, customer.Events[1].GetType())
		})
	}
}

func Test_Customer_Delete(t *testing.T) {

}

func Test_Customer_Activate(t *testing.T) {
	// TODO: Implement
}

func Test_Customer_Deactivate(t *testing.T) {
	// TODO: Implement
}
