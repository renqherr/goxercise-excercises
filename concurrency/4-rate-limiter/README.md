# Rate Limiter
## Rate Limiter

Implement a token bucket rate limiter to control the rate of operations.

## Recommended Prerequisites

Complete these exercises first for the best learning experience:

- Goroutine Counter
- Channel Ping-Pong

## Background

Rate limiting controls how frequently operations can occur. The token bucket algorithm is a popular approach:

- A bucket holds tokens (like a buffered channel)
- Tokens are added at a fixed rate
- Operations consume tokens
- If no tokens available, operations must wait or are rejected

This is essential for API rate limiting, preventing resource exhaustion, and ensuring fair resource usage.

## Your Task

Implement a rate limiter with:

1. Token bucket using channels
2. Background goroutine that replenishes tokens
3. `Allow()` - Non-blocking check for available token
4. `Wait()` - Blocking wait for a token

## Key Concepts

- **Token Bucket Algorithm**: Classic rate limiting using tokens that replenish over time
- **Burst Capacity**: Initial bucket size allows short bursts while maintaining average rate
- **Non-Blocking vs Blocking**: Allow() returns immediately, Wait() blocks for token
- **Buffered Channels as Buckets**: Channel capacity naturally implements token bucket
- **Background Replenishment**: Goroutine with ticker adds tokens at fixed rate

