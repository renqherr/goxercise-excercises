# Concurrency Exercises

## Concurrency Exercises

Practice goroutines, channels, and concurrency patterns with these exercises progressing from basic to advanced.

### Exercises in This Section
| Exercise            | Difficulty | Time   | Description                                     |
|---------------------|------------|--------|-------------------------------------------------|
| Goroutine Counter   | Easy       | 10 min | Build a thread-safe counter using goroutines    |
| Channel Ping-Pong   | Easy       | 10 min | Implement bidirectional channel communication   |
| Worker Pool         | Medium     | 20 min | Create a pool of workers processing jobs        |
| Rate Limiter        | Medium     | 25 min | Build a token bucket rate limiter               |
| Pipeline Processor  | Hard       | 30 min | Implement a concurrent data processing pipeline |

### Key Concepts Covered

- **Goroutines**: Lightweight concurrent execution
- **Channels**: Safe communication between goroutines
- **Select**: Multiplexing channel operations
- **Sync primitives**: Mutex, WaitGroup, Once
- **Patterns**: Worker pools, fan-out/fan-in, pipelines

### Prerequisites

Before starting these exercises, you should be comfortable with:

- Basic Go syntax and types
- Functions and methods
- Chapter 4: Goroutines & Channels
- Chapter 5: Sync Primitives

### Tips

1. Always use channels or sync primitives to protect shared state
2. Consider using context.Context for cancellation in real-world code
3. Use go run -race to detect race conditions

