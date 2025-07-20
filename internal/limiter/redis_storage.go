package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisStorage is a Redis-backed storage implementation for the rate limiter.
type RedisStorage struct {
	client *redis.Client
}

// NewRedisStorage creates a new RedisStorage.
func NewRedisStorage(addr, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStorage{client: rdb}
}

// Get retrieves a value from Redis.
func (s *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// Set sets a value in Redis with a TTL.
func (s *RedisStorage) Set(ctx context.Context, key string, value string, ttl int) error {
	return s.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}
