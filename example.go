package ratelimiter
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// 	import package
// )

// func main() {
// 	// Create Redis client
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	// Allow 10 requests per 30 seconds per user+IP
// 	rl := ratelimiter.New(rdb, 10, 30*time.Second)

// 	// Example protected route
// 	http.Handle("/api", rl.Middleware(http.HandlerFunc(handler)))

// 	fmt.Println("Server started on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Request successful!"))
// }




//Manual Usage Without Middleware
// ok, err := rl.Allow(req)
// if err != nil {
// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 	return
// }
// if !ok {
// 	http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
// 	return
// }
