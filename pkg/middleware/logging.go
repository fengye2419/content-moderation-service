package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function that logs incoming requests and outgoing responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("Incoming Request: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the outgoing response
		log.Printf("Outgoing Response: %s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
