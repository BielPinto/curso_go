package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"rate-limiter/internal/config"
	"rate-limiter/internal/limiter"
	"rate-limiter/internal/storage"
)

func TestRateLimitMiddleware(t *testing.T) {
	cfg := &config.Config{
		RateLimitIP: 1,
		BlockTimeIP: 5,
	}
	storage := storage.NewMockStorage() // Assuming we add MockStorage to storage package
	rl := limiter.NewRateLimiter(storage, cfg)

	handler := RateLimitMiddleware(rl)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1:1234"

	// First request
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Second request should be blocked
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status 429, got %d", rr.Code)
	}
}
