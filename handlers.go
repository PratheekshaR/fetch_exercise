package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var receiptStore = make(map[string]Receipt)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	// Parse the incoming JSON receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate a unique receipt ID
	id := generateReceiptID()

	// Store receipt in-memory
	receiptStore[id] = receipt

	// Respond with the receipt ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	// Get the receipt ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the receipt from the store
	receipt, exists := receiptStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points
	points := calculatePoints(receipt)

	// Respond with points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
