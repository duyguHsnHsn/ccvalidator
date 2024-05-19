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

	http.HandleFunc("/validate", api.HandleRequestWithWorkerPool)

	log.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	api.Wp.Close()
}
