package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"drxc00/togo/internal/models"
	"drxc00/togo/internal/services"

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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := models.Response{
			Message: "Error reading request body",
			Success: false,
			Data:    models.Togo{},
		}

		jsonResponse, err := res.Print()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, jsonResponse)
		return
	}

	var requestBody models.TogoRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := models.Response{
			Message: "Error parsing request body",
			Success: false,
			Data:    models.Togo{},
		}
		jsonResponse, err := res.Print()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, jsonResponse)
		return
	}

	togo, ok := services.CreateTogo(requestBody)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Message: "To go created successfully",
		Success: true,
		Data:    togo,
	}

	jsonResponse, err := response.Print()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResponse)
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
