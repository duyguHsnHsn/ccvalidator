package api

import (
	"ccvalidator/luhn"
	"encoding/json"
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

	isValid := luhn.Validate(req.CardNumber)
	result := ValidationResult{
		IsValid: isValid,
		Message: "Validation complete",
	}
	if !isValid {
		result.Message = "Credit card number is invalid"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
