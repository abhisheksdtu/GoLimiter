package limiter

import (
	"context"
	"sync"
	"time"
)

// MemoryStorage is an in-memory storage implementation for the rate limiter.
type MemoryStorage struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewMemoryStorage creates a new MemoryStorage.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		store: make(map[string]string),
	}
}

// Get retrieves a value from the in-memory store.
func (s *MemoryStorage) Get(ctx context.Context, key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.store[key], nil
}

// Set sets a value in the in-memory store.
func (s *MemoryStorage) Set(ctx context.Context, key string, value string, ttl int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
	if ttl > 0 {
		time.AfterFunc(time.Duration(ttl)*time.Second, func() {
			s.mu.Lock()
			delete(s.store, key)
			s.mu.Unlock()
		})
	}
	return nil
}
