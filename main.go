package main

import (
	"ccvalidator/api"
	"log"
	"net/http"
	"runtime"
)

func main() {
	jobs := make(chan api.Job, 100)

	api.StartWorkerPool(runtime.NumCPU(), jobs)

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		api.HandleRequest(jobs, w, r)
	})

	log.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
