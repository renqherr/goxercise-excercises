package main

import (
	"fmt"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID int
	Stage string
	Data  string
	Valid bool
}

// stage1 parses input jobs
func stage1(jobs <-chan Job) <-chan Result {
	// TODO: Create output channel
	// TODO: Start goroutine that reads from jobs channel
	// TODO: For each job, parse and send Result with Stage="stage1"
	// TODO: Close output channel when done
	return nil
}

// stage2 transforms data (this will be fanned out)
func stage2(input <-chan Result) <-chan Result {
	// TODO: Create output channel
	// TODO: Start goroutine that processes each result
	// TODO: Transform data (e.g., uppercase) and send with Stage="stage2"
	// TODO: Close output channel when done
	return nil
}

// stage3 validates results
func stage3(input <-chan Result) <-chan Result {
	// TODO: Create output channel
	// TODO: Start goroutine that validates each result
	// TODO: Set Valid=true if data is not empty, send with Stage="stage3"
	// TODO: Close output channel when done
	return nil
}

// merge combines multiple result channels into one
func merge(channels ...<-chan Result) <-chan Result {
	// TODO: Create output channel
	// TODO: Start WaitGroup to track goroutines
	// TODO: For each input channel, start a goroutine that forwards to output
	// TODO: Wait for all goroutines to finish, then close output
	return nil
}

func main() {
	// Create jobs
	jobs := make(chan Job, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- Job{ID: i, Data: fmt.Sprintf("item-%d", i)}
		}
		close(jobs)
	}()

	// Build pipeline
	stage1Out := stage1(jobs)

	// Fan-out to 2 workers for stage2
	stage2Out1 := stage2(stage1Out)
	stage2Out2 := stage2(stage1Out)

	// Merge stage2 outputs
	stage2Merged := merge(stage2Out1, stage2Out2)

	// Fan-out to 2 workers for stage3
	stage3Out1 := stage3(stage2Merged)
	stage3Out2 := stage3(stage2Merged)

	// Merge final outputs
	finalOut := merge(stage3Out1, stage3Out2)

	// Collect results
	for result := range finalOut {
		fmt.Printf("Job %d completed %s: %s (valid=%v)\n",
			result.JobID, result.Stage, result.Data, result.Valid)
	}
}
