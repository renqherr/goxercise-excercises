package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	// TODO: Process jobs and send results
	defer wg.Done()
}

func main() {
	numWorkers := 3
	numJobs := 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// TODO: Start workers
	var wg sync.WaitGroup

	// TODO: Send jobs

	// TODO: Close jobs channel and wait for workers

	// TODO: Collect and print results
	close(results)

	for result := range results {
		fmt.Printf("Job %d: %d\n", result.JobID, result.Output)
	}
}
