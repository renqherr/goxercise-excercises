#  Channel Ping-Pong
## Channel Ping-Pong

Practice bidirectional communication between goroutines using channels.

## Background

Channels enable goroutines to communicate and synchronize. In this exercise, you’ll implement a ping-pong pattern where two goroutines alternate sending messages back and forth.

## Your Task

Implement a ping-pong game where:

1. Two goroutines pass a ball back and forth
2. Each goroutine increments the ball’s hit count
3. The game ends after 5 total hits
4. Print each hit as it happens

## Key Concepts

- **Directional channels**: Using <-chan for receive-only and chan<- for send-only
- **Channel range**: Iterating over channel values with for range
- **Coordination**: Using a done channel to signal completion
