package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
)

// Handler handles HTTP requests for account operations
type AccountQueryHandler struct {
	accountQueryService AccountQueryService
}

// NewQueryHandler creates a new account query handler
func NewQueryHandler(accountQueryService AccountQueryService) *AccountQueryHandler {
	return &AccountQueryHandler{
		accountQueryService: accountQueryService,
	}
}

type GetAccountRequest struct {
	AccountID string
}

// GetAccount handles retrieving an account by ID
func (h *AccountQueryHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	req := &GetAccountRequest{
		AccountID: accountID,
	}

	account, err := h.accountQueryService.GetAccount(r.Context(), applicationaccount.GetAccountDTO{
		AccountID: req.AccountID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

type GetCustomerAccountsRequest struct {
	CustomerID string
}

// GetCustomerAccounts handles retrieving all accounts for a customer
func (h *AccountQueryHandler) GetCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerId"]

	req := &GetCustomerAccountsRequest{
		CustomerID: customerID,
	}

	accounts, err := h.accountQueryService.GetCustomerAccounts(r.Context(), applicationaccount.GetCustomerAccountsDTO{
		CustomerID: req.CustomerID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
