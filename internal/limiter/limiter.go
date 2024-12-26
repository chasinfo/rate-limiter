package limiter

import (
	"context"
	"errors"
	"sync"
	"time"
)

type RateLimiter struct {
	limits       map[string]int
	blockTime    time.Duration
	requests     map[string]int
	mu           sync.Mutex
	expiration   map[string]time.Time
}

func NewRateLimiter(limits map[string]int, blockTime time.Duration) *RateLimiter {
	return &RateLimiter{
		limits:     limits,
		blockTime:  blockTime,
		requests:   make(map[string]int),
		expiration: make(map[string]time.Time),
	}
}

func (rl *RateLimiter) Allow(key string) (bool, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if expiration, exists := rl.expiration[key]; exists && time.Now().Before(expiration) {
		return false, errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
	}

	limit, exists := rl.limits[key]
	if !exists {
		limit = rl.limits["default"]
	}

	if rl.requests[key] < limit {
		rl.requests[key]++
		return true, nil
	}

	rl.expiration[key] = time.Now().Add(rl.blockTime)
	return false, errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
}

func (rl *RateLimiter) Reset(key string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.requests, key)
	delete(rl.expiration, key)
}