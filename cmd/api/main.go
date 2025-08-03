package main

import (
	"drxc00/togo/internal/handler"
	"fmt"
	"net/http"

	"drxc00/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Router setup
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", middleware.Chain(
		func(w http.ResponseWriter, r *http.Request) {
			// Return a JSON response
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, `{"message": "Hello, World!"}`)
		},
		middleware.Logging,
	)).Methods("GET")

	// TOGO API Routes
	router.HandleFunc("/togo/create", middleware.Chain(
		handler.CreateTogo,
		middleware.Logging,
	)).Methods("POST")

	router.HandleFunc("/togo/get/{togo_id}", middleware.Chain(
		handler.GetTogo,
		middleware.Logging,
	)).Methods("GET")

	router.HandleFunc("/togo/update/{togo_id}", middleware.Chain(
		handler.UpdateTogo,
		middleware.Logging,
	)).Methods("PUT")

	router.HandleFunc("/togo/delete/{togo_id}", middleware.Chain(
		handler.DeleteTogo,
		middleware.Logging,
	)).Methods("DELETE")

	// Start server
	http.ListenAndServe(":8080", router)
}
