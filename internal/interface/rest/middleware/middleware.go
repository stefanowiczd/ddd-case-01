package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware represents a middleware function
type Middleware func(http.Handler) http.Handler

// Chain chains multiple middleware functions
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Logging middleware logs HTTP requests
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
