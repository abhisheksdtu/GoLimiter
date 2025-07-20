package limiter

import "context"

// Storage is the interface for rate limiter storage.
type Storage interface {
	// Get gets the value for a given key.
	Get(ctx context.Context, key string) (string, error)
	// Set sets the value for a given key with an expiration.
	Set(ctx context.Context, key string, value string, ttl int) error
}
