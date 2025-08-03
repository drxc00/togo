package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getOrderID(r *http.Request) (string, bool) {
	// Extract the order ID from the URL parameters
	vars := mux.Vars(r)
	orderID, ok := vars["togo_id"]
	if !ok {
		return "", false
	}
	return orderID, true
}

func CreateTogo(w http.ResponseWriter, r *http.Request) {
	orderID := "12345"
	response := fmt.Sprintf(`{"message": "To go created successfully", "togo_id": "%s"}`, orderID)

	// Write response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, response)
}

func GetTogo(w http.ResponseWriter, r *http.Request) {
	// Get the order ID from the URL
	orderID, ok := getOrderID(r)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := fmt.Sprintf(`{"message": "To go retrieved successfully", "togo_id": "%s"}`, orderID)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func UpdateTogo(w http.ResponseWriter, r *http.Request) {
	// Get the order ID from the URL
	orderID, ok := getOrderID(r)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Simulate updating a to-go order
	response := fmt.Sprintf(`{"message": "To go updated successfully", "togo_id": "%s"}`, orderID)

	// Write response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func DeleteTogo(w http.ResponseWriter, r *http.Request) {
	// Get the order ID from the URL
	orderID, ok := getOrderID(r)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Simulate deleting a to-go order
	response := fmt.Sprintf(`{"message": "To go deleted successfully", "order_id": "%s"}`, orderID)

	// Write response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}
