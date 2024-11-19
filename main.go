package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// POST /receipts/process
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")

	// GET /receipts/{id}/points
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")

	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

// Generate a unique receipt ID
func generateReceiptID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}
