package middleware

import (
	"log"
	"net/http"

	"rate-limiter/internal/limiter"
)

func RateLimitMiddleware(rl *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("RateLimitMiddleware: %s %s", r.Method, r.URL.Path)
			if !rl.Allow(r) {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
