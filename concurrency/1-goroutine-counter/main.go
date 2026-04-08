package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// TODO: Add fields for channel-based synchronization
}

func NewCounter() *Counter {
	// TODO: Initialize the counter
	return &Counter{}
}

func (c *Counter) Increment() {
	// TODO: Increment the counter safely
}

func (c *Counter) Value() int {
	// TODO: Return the current value safely
	return 0
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup

	// Increment from 100 goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())
}
