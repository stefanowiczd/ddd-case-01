package router

import (
	"net/http"

	accounthandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account"
)

// registerAccountRoutes registers all account-related routes
func RegisterAccountRoutes(r *http.ServeMux, aqh *accounthandler.AccountQueryHandler, ah *accounthandler.AccountHandler) {

	// Query operations:
	// Get account / accounts
	r.HandleFunc("GET /accounts/{id}", aqh.GetAccount)
	r.HandleFunc("GET /customers/{customerId}/accounts", aqh.GetCustomerAccounts)

	// Mutate operations:
	// Create
	r.HandleFunc("POST /account", ah.CreateAccount)

	// Block / unblock
	r.HandleFunc("POST /accounts/{id}/block", ah.BlockAccount)
	r.HandleFunc("POST /accounts/{id}/unblock", ah.UnblockAccount)

	// Deposit / withdrawn
	r.HandleFunc("POST /accounts/{id}/deposit", ah.Deposit)
	r.HandleFunc("POST /accounts/{id}/withdraw", ah.Withdraw)

}
