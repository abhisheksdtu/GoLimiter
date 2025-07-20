package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration.
type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	ServerPort    string
	RateLimit     int
	Burst         int
}

// Load loads configuration from environment variables.
func Load() *Config {
	return &Config{
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		RateLimit:     getEnvAsInt("RATE_LIMIT", 1), // Default to 10 requests per second
		Burst:         getEnvAsInt("BURST", 1),      // Default to a burst of 5
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return fallback
}
