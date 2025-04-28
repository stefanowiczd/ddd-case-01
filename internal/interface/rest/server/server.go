package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/middleware"
	"github.com/stefanowiczd/ddd-case-01/internal/interface/rest/router"

	accounthandler "github.com/stefanowiczd/ddd-case-01/internal/interface/rest/handler/account"
)

// Server represents the HTTP server
type Server struct {
	router *http.ServeMux
	server *http.Server
}

// Config holds the server configuration
type Config struct {
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// NewServer creates a new HTTP server
func NewServer(
	config Config,
	accountQueryHandler *accounthandler.AccountQueryHandler,
	accountHandler *accounthandler.AccountHandler,
) *Server {
	// Create router
	r := http.NewServeMux()

	// Apply middleware
	handler := middleware.Chain(
		r,
		middleware.Logging,
	)

	// Register routes
	router.RegisterAccountRoutes(r, accountQueryHandler, accountHandler)

	// Create server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	return &Server{
		router: r,
		server: srv,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// DefaultConfig returns the default server configuration
func DefaultConfig() Config {
	return Config{
		Port:            8080,
		ReadTimeout:     15 * time.Second,
		WriteTimeout:    15 * time.Second,
		ShutdownTimeout: 5 * time.Second,
	}
}
