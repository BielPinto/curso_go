package limiter

import (
	"net/http"
	"testing"

	"rate-limiter/internal/config"
	"rate-limiter/internal/storage"
)

func TestRateLimiter_Allow_IP(t *testing.T) {
	cfg := &config.Config{
		RateLimitIP:    2,
		BlockTimeIP:    5,
		RateLimitToken: 10,
		BlockTimeToken: 5,
	}
	storage := storage.NewMockStorage()
	rl := NewRateLimiter(storage, cfg)

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1:1234"

	// First request
	if !rl.Allow(req) {
		t.Error("First request should be allowed")
	}

	// Second request
	if !rl.Allow(req) {
		t.Error("Second request should be allowed")
	}

	// Third request should be blocked
	if rl.Allow(req) {
		t.Error("Third request should be blocked")
	}
}

func TestRateLimiter_Allow_Token(t *testing.T) {
	cfg := &config.Config{
		RateLimitIP:    1,
		BlockTimeIP:    5,
		RateLimitToken: 2,
		BlockTimeToken: 5,
	}
	storage := storage.NewMockStorage()
	rl := NewRateLimiter(storage, cfg)

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1:1234"
	req.Header.Set("API_KEY", "token123")

	// First request
	if !rl.Allow(req) {
		t.Error("First request should be allowed")
	}

	// Second request
	if !rl.Allow(req) {
		t.Error("Second request should be allowed")
	}

	// Third request should be blocked
	if rl.Allow(req) {
		t.Error("Third request should be blocked")
	}
}
