package ratelimiter

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	RedisClient *redis.Client
	Limit       int           // Number of allowed requests in the duration window
	Window      time.Duration // Sliding window duration
}


// New will create a new instance of RateLimiter
func New(redisClient *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		RedisClient: redisClient,
		Limit:       limit,
		Window:      window,
	}
}
