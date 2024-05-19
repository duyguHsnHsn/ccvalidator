// api/rate_limiter.go

package api

import (
	"net/http"
	"sync"
	"time"
)

var (
	requestCounts = make(map[string]int)
	mutex         sync.Mutex
)

const (
	limit     = 5
	resetTime = time.Minute
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mutex.Lock()
		count := requestCounts[ip]
		if count >= limit {
			mutex.Unlock()
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		requestCounts[ip] = count + 1
		mutex.Unlock()
		time.AfterFunc(resetTime, func() {
			mutex.Lock()
			requestCounts[ip]--
			mutex.Unlock()
		})
		next.ServeHTTP(w, r)
	})
}
