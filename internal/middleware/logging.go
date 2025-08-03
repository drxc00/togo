package middleware

import (
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Logging Middleware that logs the request method, URL, and response status code
func Logging(n http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		n.ServeHTTP(lrw, r)
		log.Printf("Response status code: %d", lrw.statusCode)
	})
}
