package customer

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	customerapplication "github.com/stefanowiczd/ddd-case-01/internal/application/customer"
)

// CustomerQueryHandler handles HTTP requests for customer query operations
type CustomerQueryHandler struct {
	customerService CustomerQueryService
}

func NewCustomerQueryHandler(customerService CustomerQueryService) *CustomerQueryHandler {
	return &CustomerQueryHandler{
		customerService: customerService,
	}
}

func (h *CustomerQueryHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerId")

	if _, err := uuid.Parse(customerID); err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.GetCustomer(
		r.Context(),
		customerapplication.GetCustomerDTO{
			CustomerID: customerID,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(customer) // TODO decide about handling of this error.
}
