// main.go

package main

import (
	"ccvalidator/api"
	"log"
	"net/http"
)

func main() {
	workerCount := 5
	api.InitWorkerPool(workerCount)

	mux := http.NewServeMux()
	mux.HandleFunc("/validate", api.HandleRequestWithWorkerPool)

	rateLimitedMux := api.RateLimiter(mux)

	log.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", rateLimitedMux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	api.Wp.Close()
}
