package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// receiptMap represents the in-memory database
var receiptMap map[string]Receipt

// Could contain other complex initializations too
func init() {
	receiptMap = make(map[string]Receipt)
}

type Handler struct {
}

// Method in Handler to register the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/process", h.processReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", h.getReceiptPointsHandler).Methods("GET")

}

// Handler for POST /receipts/process
func (h *Handler) processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt

	// Decode JSON body into Receipt struct
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) == 0 {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Define the pattern for the Retailer field
	retailerPattern := `^[\w\s\-&]+$`
	retailerRegex, err := regexp.Compile(retailerPattern)
	if err != nil {
		http.Error(w, "Error compiling retailer regex", http.StatusInternalServerError)
		return
	}
	if !retailerRegex.MatchString(receipt.Retailer) {
		http.Error(w, "The retailer field is invalid", http.StatusBadRequest)
		return
	}

	// Define the pattern for the Total field
	totalPattern := `^\d+\.\d{2}$`
	totalRegex, err := regexp.Compile(totalPattern)
	if err != nil {
		http.Error(w, "Error compiling total regex", http.StatusInternalServerError)
		return
	}
	if !totalRegex.MatchString(receipt.Total) {
		http.Error(w, "The total field is invalid", http.StatusBadRequest)
		return
	}

	// Validate purchaseDate format (must be "2006-01-02")
	_, err = time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		http.Error(w, "The purchaseDate field must be in YYYY-MM-DD format", http.StatusBadRequest)
		return
	}

	// Validate purchaseTime format
	_, err = time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		http.Error(w, "The purchaseTime field must be in the valid HH:MM format", http.StatusBadRequest)
		return
	}

	// Define the pattern for the shortDescription field in items
	shortDescriptionPattern := `^[\w\s\-]+$`
	shortDescRegex, err := regexp.Compile(shortDescriptionPattern)
	if err != nil {
		http.Error(w, "Error compiling short description regex", http.StatusInternalServerError)
		return
	}

	for _, item := range receipt.Items {
		if item.ShortDescription == "" || item.Price == "" {
			http.Error(w, "Each item must have a shortDescription and a price", http.StatusBadRequest)
			return
		}
		if !shortDescRegex.MatchString(item.ShortDescription) {
			http.Error(w, "The shortDescription field in items is invalid", http.StatusBadRequest)
			return
		}
		if !totalRegex.MatchString(item.Price) {
			http.Error(w, "The Price field is invalid", http.StatusBadRequest)
			return
		}
	}

	// Generate unique id
	id := uuid.New().String()

	// Store receipt in map (assuming receiptMap is defined somewhere)
	receiptMap[id] = receipt

	// Respond back with the generated id for the receipt
	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for GET /receipts/{id}/points
func (h *Handler) getReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve receipt from map
	receipt, found := receiptMap[id]
	if !found {
		http.Error(w, "No receipt found for that id", http.StatusNotFound)
		return
	}

	// Calculate points based on receipt rules
	points := calculatePoints(receipt)

	// Respond back with points
	response := ReceiptPoints{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Return the handler pointer
func NewHandler() *Handler {
	return &Handler{}
}
