package handlers

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu           sync.Mutex
	refillPeriod time.Time
	tokenRefill  int
	tokens       int
	maxTokens    int
}

func NewRateLimiter(maxTokens, refillRate int, refillPeriod time.Time) *RateLimiter {
	return &RateLimiter{
		tokens:       maxTokens,
		maxTokens:    maxTokens,
		tokenRefill:  refillRate,
		refillPeriod: refillPeriod,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Refill tokens if needed
	elapsed := time.Since(rl.refillPeriod)
	tokensToAdd := int(elapsed.Seconds()) * rl.tokenRefill
	if tokensToAdd > 0 {
		rl.tokens = min(rl.tokens+tokensToAdd, rl.maxTokens)
	}

	// Check if there are enough tokens to allow the request
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
