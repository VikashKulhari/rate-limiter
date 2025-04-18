package ratelimiter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)


//Limiter can be used as middleware 
func (r *RateLimiter) Limiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		allowed, err := r.Allow(req)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !allowed {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, req)
	})
}

//Allow will check if the user is allowed to send request more requests 
func (r *RateLimiter) Allow(req *http.Request) (bool, error) {
	ctx := context.Background()

	claims, err := ExtractClaimsFromJWT(req)
	if err != nil {
		return false, err
	}

	ip := GetIPAddress(req)
	key := fmt.Sprintf("rl:%s:%s", claims.UserID, ip)

	now := time.Now().Unix()
	windowStart := now - int64(r.Window.Seconds())

	// Remove old entries outside the sliding window
	r.RedisClient.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

	// Count requests in window
	count, _ := r.RedisClient.ZCard(ctx, key).Result()
	if int(count) >= r.Limit {
		return false, nil
	}

	// Add current request timestamp
	r.RedisClient.ZAdd(ctx, key, &redis.Z{Score: float64(now), Member: fmt.Sprintf("%d", now)}).Err()

	// Set expiry for key
	r.RedisClient.Expire(ctx, key, r.Window*2)

	return true, nil
}
