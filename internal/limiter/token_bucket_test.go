package limiter

import (
	"testing"
	"time"
)

func TestTokenBucket_Allow(t *testing.T) {
	storage := NewMemoryStorage()
	tb := NewTokenBucket(storage, 2, 2) // 2 tokens/sec, burst 2

	key := "test-key"

	// First two requests should be allowed (burst)
	allowed, _ := tb.Allow(nil, key)
	if !allowed {
		t.Error("expected first request to be allowed")
	}
	allowed, _ = tb.Allow(nil, key)
	if !allowed {
		t.Error("expected second request to be allowed")
	}

	// Third request should be denied (burst exhausted)
	allowed, _ = tb.Allow(nil, key)
	if allowed {
		t.Error("expected third request to be denied")
	}

	// Wait for tokens to refill
	time.Sleep(600 * time.Millisecond)
	allowed, _ = tb.Allow(nil, key)
	if !allowed {
		t.Error("expected request to be allowed after refill")
	}
}
