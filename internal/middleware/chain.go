package middleware

import (
	"net/http"
)

// Function to create a chain of middlewares
// This makes it easy to add new middlewares to every route handler
func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// Apply the first middleware to the handler
	// The order of middlewares is reversed so that the first one in the list is applied last
	// This allows for the last middleware to be the first one executed
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	// Return the final handler with all middlewares applied
	return handler
}
