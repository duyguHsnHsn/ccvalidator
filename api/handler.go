package api

import (
	"ccvalidator/luhn"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CreditCardRequest struct {
	CardNumber string `json:"card_number"`
}

type ValidationResult struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	log.Printf("Body: %v\n", string(body))

	var req CreditCardRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("Error decoding JSON request: %v\n", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Processing validation for card number: %s\n", req.CardNumber)
	isValid := luhn.Validate(req.CardNumber)
	result := ValidationResult{
		IsValid: isValid,
		Message: "Validation complete",
	}
	if !isValid {
		result.Message = "Credit card number is invalid"
		log.Printf("Validation failed for card number: %s\n", req.CardNumber)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
	log.Printf("Response successfully written for card number: %s\n", req.CardNumber)
}
