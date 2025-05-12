package customer

import (
	"errors"
)

// Customer errors
var (
	// ErrCustomerNotFound is returned when an customer is not found.
	ErrCustomerNotFound = errors.New("customer not found")
	// ErrCustomerAlreadyExists is returned when a customer already exists.
	ErrCustomerAlreadyExists = errors.New("customer already exists")
)
