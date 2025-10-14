package limiter

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"rate-limiter/internal/config"
	"rate-limiter/internal/storage"
)

type RateLimiter struct {
	storage storage.Storage
	config  *config.Config
}

func NewRateLimiter(storage storage.Storage, config *config.Config) *RateLimiter {
	return &RateLimiter{
		storage: storage,
		config:  config,
	}
}

func (rl *RateLimiter) Allow(r *http.Request) bool {
	ctx := context.Background()

	// Extract token from header
	token := r.Header.Get("API_KEY")
	ip := getIP(r)

	var key, blockKey string
	var limit int
	var blockTime time.Duration

	if token != "" {
		key = fmt.Sprintf("token:%s", token)
		blockKey = fmt.Sprintf("block:token:%s", token)
		limit = rl.config.RateLimitToken
		blockTime = time.Duration(rl.config.BlockTimeToken) * time.Second
	} else {
		key = fmt.Sprintf("ip:%s", ip)
		blockKey = fmt.Sprintf("block:ip:%s", ip)
		limit = rl.config.RateLimitIP
		blockTime = time.Duration(rl.config.BlockTimeIP) * time.Second
	}
	log.Printf("limit %v, blockTime %v key%v", limit, blockTime, key)

	// Check if blocked
	if _, err := rl.storage.Get(ctx, blockKey); err == nil {
		return false // blocked
	}

	// Increment count
	count, err := rl.storage.Increment(ctx, key)
	if err != nil {
		log.Printf("Error incrementing key %s: %v", key, err)
		return false // error, deny
	}

	// Set expiration if first time
	if count == 1 {
		log.Printf("Setting expiration for key %s", key)
		rl.storage.SetExpiration(ctx, key, time.Second)
	}

	if count > int64(limit) {
		log.Printf("Rate limit exceeded for key %s: %d > %d", key, count, limit)
		// Exceeded, block
		rl.storage.SetExpiration(ctx, blockKey, blockTime)
		return false
	}

	return true
}

func getIP(r *http.Request) string {
	// Get IP from X-Forwarded-For or RemoteAddr
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	// Remove port if present
	if strings.Contains(ip, ":") {
		ip, _, _ = strings.Cut(ip, ":")
	}
	return ip
}
