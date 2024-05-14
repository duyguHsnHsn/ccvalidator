package main

import (
	"ccvalidator/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/validate", api.HandleRequest)

	log.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
