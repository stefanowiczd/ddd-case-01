package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
)

// Handler handles HTTP requests for account operations
type AccountHandler struct {
	accountService AccountService
}

// NewHandler creates a new account handler
func NewHandler(accountService AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// CreateAccountRequest represents the request body for creating an account
type CreateAccountRequest struct {
	CustomerID     string  `json:"customerId"`
	InitialBalance float64 `json:"initialBalance"`
	Currency       string  `json:"currency"`
}

// CreateAccount handles the creation of a new account
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	account, err := h.accountService.CreateAccount(r.Context(), applicationaccount.CreateAccountDTO{
		CustomerID:     req.CustomerID,
		InitialBalance: req.InitialBalance,
		Currency:       req.Currency,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// DepositRequest represents the request body for depositing money
type DepositRequest struct {
	Amount float64 `json:"amount"`
}

// Deposit handles depositing money into an account
func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	var req DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.accountService.Deposit(r.Context(), applicationaccount.DepositDTO{
		AccountID: accountID,
		Amount:    req.Amount,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// WithdrawRequest represents the request body for withdrawing money
type WithdrawRequest struct {
	Amount float64 `json:"amount"`
}

// Withdraw handles withdrawing money from an account
func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	var req WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.accountService.Withdraw(r.Context(), applicationaccount.WithdrawDTO{
		AccountID: accountID,
		Amount:    req.Amount,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// BlockAccount handles blocking an account
func (h *AccountHandler) BlockAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	if err := h.accountService.BlockAccount(r.Context(), applicationaccount.BlockAccountDTO{
		AccountID: accountID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UnblockAccount handles unblocking an account
func (h *AccountHandler) UnblockAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]

	if err := h.accountService.UnblockAccount(r.Context(), applicationaccount.UnblockAccountDTO{
		AccountID: accountID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
