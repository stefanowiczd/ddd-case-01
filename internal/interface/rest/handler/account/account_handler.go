package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

func (r CreateAccountRequest) Validate() error {
	if _, err := uuid.Parse(r.CustomerID); err != nil {
		return fmt.Errorf("validate: customer id ass uuid: %w", err)
	}

	return nil
}

// CreateAccount handles the creation of a new account
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	Amount    float64 `json:"amount"`
	AccountID string
}

func (r *DepositRequest) Validate() error {
	if _, err := uuid.Parse(r.AccountID); err != nil {
		return fmt.Errorf("validate: account id as uuid: %w", err)
	}

	return nil
}

// Deposit handles depositing money into an account
func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("id")

	var req DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.AccountID = accountID

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.accountService.Deposit(r.Context(), applicationaccount.DepositDTO{
		AccountID: uuid.MustParse(accountID),
		Amount:    req.Amount,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// WithdrawRequest represents the request body for withdrawing money
type WithdrawRequest struct {
	Amount    float64 `json:"amount"`
	AccountID string
}

func (r *WithdrawRequest) Validate() error {
	if _, err := uuid.Parse(r.AccountID); err != nil {
		return fmt.Errorf("validate: account id as uuid: %w", err)
	}

	return nil
}

// Withdraw handles withdrawing money from an account
func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("id")

	var req WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.AccountID = accountID

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.accountService.Withdraw(r.Context(), applicationaccount.WithdrawDTO{
		AccountID: uuid.MustParse(accountID),
		Amount:    req.Amount,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// BlockAccountRequest represents the request body for blocking an account
type BlockAccountRequest struct {
	AccountID string
}

func (r *BlockAccountRequest) Validate() error {
	if _, err := uuid.Parse(r.AccountID); err != nil {
		return fmt.Errorf("validate: account id as uuid: %w", err)
	}

	return nil
}

// BlockAccount handles blocking an account
func (h *AccountHandler) BlockAccount(w http.ResponseWriter, r *http.Request) {
	req := BlockAccountRequest{
		AccountID: r.PathValue("id"),
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.accountService.BlockAccount(r.Context(), applicationaccount.BlockAccountDTO{
		AccountID: uuid.MustParse(req.AccountID),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UnblockAccountRequest represents the request body for unblocking an account
type UnblockAccountRequest struct {
	AccountID string
}

func (r *UnblockAccountRequest) Validate() error {
	if _, err := uuid.Parse(r.AccountID); err != nil {
		return fmt.Errorf("validate: account id as uuid: %w", err)
	}

	return nil
}

// UnblockAccount handles unblocking an account
func (h *AccountHandler) UnblockAccount(w http.ResponseWriter, r *http.Request) {
	req := UnblockAccountRequest{
		AccountID: r.PathValue("id"),
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.accountService.UnblockAccount(r.Context(), applicationaccount.UnblockAccountDTO{
		AccountID: uuid.MustParse(req.AccountID),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
