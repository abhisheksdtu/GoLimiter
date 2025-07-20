package limiter

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TokenBucket implements the token bucket algorithm.
type TokenBucket struct {
	storage   Storage
	rate      int
	burst     int
	keyPrefix string
}

// NewTokenBucket creates a new TokenBucket.
func NewTokenBucket(storage Storage, rate, burst int) *TokenBucket {
	return &TokenBucket{
		storage:   storage,
		rate:      rate,
		burst:     burst,
		keyPrefix: "token_bucket",
	}
}

// Allow checks if a request is allowed for a given key.
func (tb *TokenBucket) Allow(ctx context.Context, key string) (bool, error) {
	redisKey := fmt.Sprintf("%s:%s", tb.keyPrefix, key)

	// Get the current state from storage
	val, err := tb.storage.Get(ctx, redisKey)
	if err != nil {
		return false, err
	}

	var tokens float64
	var lastRefill int64

	now := time.Now().UnixNano()

	if val == "" {
		// First request
		tokens = float64(tb.burst) - 1
		lastRefill = now
	} else {
		parts := strings.Split(val, ":")
		if len(parts) != 2 {
			return false, fmt.Errorf("invalid token bucket value in storage")
		}
		tokens, err = strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return false, err
		}
		lastRefill, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return false, err
		}

		// Refill tokens
		elapsed := float64(now-lastRefill) / float64(time.Second)
		tokens += elapsed * float64(tb.rate)
		if tokens > float64(tb.burst) {
			tokens = float64(tb.burst)
		}
		lastRefill = now
	}

	if tokens >= 1 {
		tokens--
		newValue := fmt.Sprintf("%f:%d", tokens, lastRefill)
		// TTL of burst/rate + 1 to be safe
		ttl := int(float64(tb.burst)/float64(tb.rate)) + 1
		err := tb.storage.Set(ctx, redisKey, newValue, ttl)
		return true, err
	}

	return false, nil
}
