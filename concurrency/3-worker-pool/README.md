# Worker Pool
## Worker Pool

Implement a worker pool pattern to process multiple jobs concurrently.

## Recommended Prerequisites

Complete these exercises first for the best learning experience:

- Goroutine Counter
- Channel Ping-Pong

## Background

The worker pool pattern is essential for controlling concurrency. Instead of spawning unlimited goroutines, you create a fixed number of workers that process jobs from a queue.

## Your Task

Implement a worker pool that:

1. Creates a configurable number of workers
2. Processes jobs from a shared channel
3. Each job is a number that should be squared
4. Collects results and prints them

## Key Concepts

- **Worker pool pattern**: Fixed number of workers processing from a queue
- **Buffered channels**: Using buffer to decouple producers and consumers
- **WaitGroup**: Waiting for multiple goroutines to complete
- **Channel closing**: Signaling no more values will be sent

