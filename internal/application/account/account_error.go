package account

import (
	"errors"
)

// Account errors
var (
	// ErrAccountNotFound is returned when an account is not found.
	ErrAccountNotFound = errors.New("account not found")
	// ErrCustomerNotFound is returned when a customer is not found.
	ErrCustomerNotFound = errors.New("customer not found")
	// ErrInvalidWithdrawAmount is returned when the withdraw money amount is invalid.
	ErrInvalidWithdrawAmount = errors.New("invalid withdraw money amount")
	// ErrInvalidDepositAmount is returned when the deposit money amount is invalid.
	ErrInvalidDepositAmount = errors.New("invalid deposit money amount")
	// ErrInvalidInitialBalanceAmount is returned when the initial account balance amount is invalid.
	ErrInvalidInitialBalanceAmount = errors.New("invalid initial account balance amount")
)
