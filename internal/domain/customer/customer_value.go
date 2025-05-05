package customer

// CustomerType represents the type of a customer
type CustomerType string

const (
	CustomerTypeIndividual CustomerType = "individual"
	CustomerTypeBusiness   CustomerType = "business"
)

// CustomerStatus represents the status of a customer
type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
	CustomerStatusBlocked  CustomerStatus = "blocked"
)

func (a CustomerStatus) String() string {
	return string(a)
}

// Address represents a physical address
type Address struct {
	Street     string `json:"street"`     // Street name and number
	City       string `json:"city"`       // City name
	State      string `json:"state"`      // State or province
	PostalCode string `json:"postalCode"` // Postal or ZIP code
	Country    string `json:"country"`    // Country name
}

func (a Address) compare(b Address) bool { //nolint:unused
	return a.Street == b.Street &&
		a.City == b.City &&
		a.State == b.State &&
		a.PostalCode == b.PostalCode &&
		a.Country == b.Country
}
