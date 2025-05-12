package account

import (
	"errors"
)

// Account errors
var (
	// ErrInsufficientFunds is returned when a customer has insufficient funds
	ErrInsufficientFunds = errors.New("insufficient funds")
	// ErrAccountNotFound is returned when an account is not found
	ErrAccountNotFound = errors.New("account not found")
	// ErrAccountAlreadyExists is returned when an account already exists
	ErrAccountAlreadyExists = errors.New("account already exists")
)

// Account Event errors
var (
	// ErrAccountEventNotFound is returned when an account event is not found
	ErrAccountEventNotFound = errors.New("account event not found")
)
