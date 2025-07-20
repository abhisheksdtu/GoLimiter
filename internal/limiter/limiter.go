package limiter

import "context"

// Strategy is the interface for a rate limiting algorithm.
type Strategy interface {
	// Allow checks if a request is allowed for a given key.
	Allow(ctx context.Context, key string) (bool, error)
}

// RateLimiter is the main rate limiter.
type RateLimiter struct {
	strategy Strategy
}

// NewRateLimiter creates a new RateLimiter with the given strategy.
func NewRateLimiter(strategy Strategy) *RateLimiter {
	return &RateLimiter{strategy: strategy}
}

// Allow checks if a request is allowed for a given key.
func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	return rl.strategy.Allow(ctx, key)
}
