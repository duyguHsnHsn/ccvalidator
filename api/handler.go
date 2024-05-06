package api

import (
	"net/http"
)

// HandleRequest creates a job from the HTTP request and adds it to the job channel
func HandleRequest(jobs chan<- Job, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	job := Job{ResponseWriter: w, Request: r}
	jobs <- job
}
