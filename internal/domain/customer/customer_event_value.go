package customer

// CustomerEventType represents the type of customer event
type CustomerEventType string

// String returns the string representation of the customer event type
func (e CustomerEventType) String() string {
	return string(e)
}

const (
	CustomerCreatedEventType CustomerEventType = "customer.created"
	CustomerDeletedEventType CustomerEventType = "customer.deleted"

	CustomerActivatedEventType   CustomerEventType = "customer.activated"
	CustomerDeactivatedEventType CustomerEventType = "customer.deactivated"

	CustomerBlockedEventType   CustomerEventType = "customer.blocked"
	CustomerUnblockedEventType CustomerEventType = "customer.unblocked"

	CustomerUpdatedNameEventType    CustomerEventType = "customer.updated.name"
	CustomerUpdatedContactEventType CustomerEventType = "customer.updated.contact"
	CustomerUpdatedAddressEventType CustomerEventType = "customer.updated.address"
	CustomerUpdatedAllEventType     CustomerEventType = "customer.updated.all"
)
