# Goroutine Counter
## Goroutine Counter

Build a thread-safe counter that can be safely incremented from multiple goroutines.

## Background

When multiple goroutines access shared state concurrently, you need synchronization to prevent race conditions. In this exercise, you’ll build a counter that uses a channel to serialize access.

## Your Task

Implement a Counter struct with:

1. A method to increment the counter
2. A method to get the current value
3. Proper synchronization using channels

The counter should work correctly even when incremented from 100 concurrent goroutines.

## Key Concepts

- **Channel-based synchronization**: Using channels to serialize access to shared state
- **Select statement**: Handling multiple channel operations
- **Ownership pattern**: One goroutine “owns” the state and communicates via channels
