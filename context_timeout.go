package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Simulate a long-running task
		fmt.Println("Task completed successfully")
	case <-ctx.Done(): // Handle cancellation or timeout
		fmt.Println("Task canceled or timed out:", ctx.Err())
	}
}

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure the context is canceled to release resources

	// Start the long-running task in a goroutine
	go longRunningTask(ctx)

	// Wait for the task to finish or timeout
	time.Sleep(3 * time.Second) // Wait longer than the timeout to observe the effect
}
