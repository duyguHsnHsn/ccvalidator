package api

import (
	"net/http"
)

// HandleRequest creates a job from the HTTP request and adds it to the job channel
func HandleRequest(jobs chan<- Job, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	job := Job{ResponseWriter: w, Request: r}
	jobs <- job
}
