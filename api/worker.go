package api

import (
	"ccvalidator/luhn"
	"encoding/json"
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

type Job struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func StartWorkerPool(numWorkers int, jobs chan Job) {
	for i := 0; i < numWorkers; i++ {
		go worker(jobs)
	}
}

func worker(jobs chan Job) {
	for job := range jobs {
		processRequest(job.ResponseWriter, job.Request)
	}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreditCardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
	}
}
