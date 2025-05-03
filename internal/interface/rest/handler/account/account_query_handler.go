package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
)

// AccountQueryHandler handles HTTP requests for account query operations
type AccountQueryHandler struct {
	accountQueryService AccountQueryService
}

// NewAccountQueryHandler creates a new account query handler
func NewAccountQueryHandler(accountQueryService AccountQueryService) *AccountQueryHandler {
	return &AccountQueryHandler{
		accountQueryService: accountQueryService,
	}
}

type GetAccountRequest struct {
	AccountID string
}

func (r GetAccountRequest) Validate() error {
	if _, err := uuid.Parse(r.AccountID); err != nil {
		return fmt.Errorf("validate: account id as uuid: %w", err)
	}

	return nil
}

// GetAccount handles retrieving an account by ID
func (h *AccountQueryHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	req := &GetAccountRequest{
		AccountID: r.PathValue("id"),
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountQueryService.GetAccount(r.Context(), applicationaccount.GetAccountDTO{
		AccountID: uuid.MustParse(req.AccountID),
	})
	if err != nil {
		// TODO add more detailed error validation
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(account) // TODO decide about handling of this error.
}

type GetCustomerAccountsRequest struct {
	CustomerID string
}

func (r GetCustomerAccountsRequest) Validate() error {
	if _, err := uuid.Parse(r.CustomerID); err != nil {
		return fmt.Errorf("validate: customer id as uuid: %w", err)
	}

	return nil
}

// GetCustomerAccounts handles retrieving all accounts for a customer
func (h *AccountQueryHandler) GetCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	req := &GetCustomerAccountsRequest{
		CustomerID: r.PathValue("customerId"),
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accounts, err := h.accountQueryService.GetCustomerAccounts(r.Context(), applicationaccount.GetCustomerAccountsDTO{
		CustomerID: uuid.MustParse(req.CustomerID),
	})
	if err != nil {
		// TODO add more detailed error validation
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(accounts) // TODO decide about handling of this error.
}
