package customer

import (
	"errors"
)

// Customer errors
var (
	// ErrCustomerAlreadyExists is returned when a customer already exists
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	// ErrCustomerNotFound is returned when a customer is not found
	ErrCustomerNotFound = errors.New("customer not found")
)

// Customer Event errors
var (
	// ErrCustomerEventNotFound is returned when a customer event is not found
	ErrCustomerEventNotFound = errors.New("customer event not found")
)
