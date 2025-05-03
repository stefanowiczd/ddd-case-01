package main

import (
	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
	applicationcustomer "github.com/stefanowiczd/ddd-case-01/internal/application/customer"
	accounthandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account"
	customerhandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/customer"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/server"
)

func main() {

	accountQueryService := &applicationaccount.AccountService{}
	customerQueryService := &applicationcustomer.CustomerService{}

	accountService := &applicationaccount.AccountService{}
	customerService := &applicationcustomer.CustomerService{}

	accountQueryHandler := accounthandler.NewQueryHandler(
		accountQueryService,
	)

	accountHandler := accounthandler.NewHandler(
		accountService,
	)

	customerQueryHandler := customerhandler.NewCustomerQueryHandler(
		customerQueryService,
	)

	customerHandler := customerhandler.NewCustomerHandler(
		customerService,
	)

	server := server.NewServer(
		server.DefaultConfig(),
		accountQueryHandler,
		customerQueryHandler,
		accountHandler,
		customerHandler,
	)
	_ = server.Start() // TODO decide about handling of this error.
}
