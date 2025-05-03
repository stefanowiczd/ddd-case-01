package router

import (
	"net/http"

	customerhandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/customer"
)

// registerCustomerRoutes registers all customer-related routes
func RegisterCustomerRoutes(
	r *http.ServeMux,
	cqh *customerhandler.CustomerQueryHandler,
	ch *customerhandler.CustomerHandler,
) {

	// Query operations:
	// Get account / accounts
	r.HandleFunc("GET /customers/{customerId}", cqh.GetCustomer)

	// Mutate operations:
	// Create
	r.HandleFunc("POST /customers", ch.CreateCustomer)

	// Block / unblock
	r.HandleFunc("POST /customers/{customerId}/block", ch.BlockCustomer)
	r.HandleFunc("POST /customers/{customerId}/unblock", ch.UnblockCustomer)

	// Update
	r.HandleFunc("PUT /customers/{customerId}", ch.UpdateCustomer)

	// Delete
	r.HandleFunc("DELETE /customers/{customerId}", ch.DeleteCustomer)

}
