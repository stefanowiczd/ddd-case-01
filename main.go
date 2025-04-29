package main

import (
	applicationaccount "github.com/stefanowiczd/ddd-case-01/internal/application/account"
	accounthandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/server"
)

func main() {

	accountQueryService := &applicationaccount.AccountService{}

	accountService := &applicationaccount.AccountService{}

	accountQueryHandler := accounthandler.NewQueryHandler(
		accountQueryService,
	)

	accountHandler := accounthandler.NewHandler(
		accountService,
	)

	server := server.NewServer(server.DefaultConfig(), accountQueryHandler, accountHandler)
	_ = server.Start() // TODO decide about handling of this error.
}
