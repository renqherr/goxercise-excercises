package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens   chan struct{}
	rate     time.Duration // Time between token additions
	capacity int           // Maximum tokens
}

func NewRateLimiter(tokensPerSecond int, capacity int) *RateLimiter {
	rl := &RateLimiter{
		// TODO: Create buffered channel with size = capacity
		rate:     time.Second / time.Duration(tokensPerSecond),
		capacity: capacity,
	}

	// TODO: Fill the bucket initially
	// Add 'capacity' tokens to the channel

	// TODO: Start goroutine to replenish tokens
	// Use time.NewTicker with rl.rate interval

	return rl
}

func (rl *RateLimiter) Allow() bool {
	// TODO: Try to get a token without blocking
	// Use select with default case
	// Return true if token acquired, false otherwise
	return false
}

func (rl *RateLimiter) Wait() {
	// TODO: Block until a token is available
	// Simply receive from the tokens channel
}

func main() {
	// Create rate limiter: 5 tokens per second, burst of 5
	limiter := NewRateLimiter(5, 5)

	// Test immediate burst (should succeed 5 times)
	fmt.Println("Testing burst capacity...")
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: allowed\n", i+1)
		}
	}

	// 6th request should be rate limited
	if !limiter.Allow() {
		fmt.Println("Request 6: rate limited (expected)")
	}

	// Wait for token replenishment
	fmt.Println("\nWaiting for token...")
	limiter.Wait()
	fmt.Println("Token available, request allowed")

	// Test sustained rate
	fmt.Println("\nTesting sustained rate...")
	start := time.Now()
	for i := 0; i < 3; i++ {
		limiter.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Request %d at %v\n", i+1, elapsed.Round(time.Millisecond))
	}
}
