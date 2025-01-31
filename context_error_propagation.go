package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Generate initial number
	num := rand.Intn(100)
	fmt.Println("task1: Initial number", num)

	// First sleep phase (3 seconds)
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("task1: Completed first sleep")
	case <-ctx.Done():
		fmt.Println("task1: Canceled during first sleep")
		return
	}

	// Add 1 to the number
	num++
	fmt.Println("task1: After addition", num)

	// Second sleep phase (2 seconds)
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("task1: Completed second sleep")
		fmt.Println("task1: Final result", num)
	case <-ctx.Done():
		fmt.Println("task1: Canceled during second sleep")
		return
	}
}

func task2(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()

	// Generate random number (1-10)
	num := rand.Intn(10) + 1
	fmt.Println("task2: Generated number", num)

	// Check for even number
	if num%2 == 0 {
		fmt.Println("task2: Even number detected, canceling context!")
		cancel()
	}

	// First sleep phase (2 seconds)
	fmt.Println("task2: Starting first sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("task2: Completed first sleep")

	// Subtract 1 from the number
	num--
	fmt.Println("task2: After subtraction", num)

	// Second sleep phase (1 second)
	fmt.Println("task2: Starting second sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("task2: Completed second sleep")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(2)
	go task1(ctx, &wg)
	go task2(ctx, &wg, cancel)

	wg.Wait()
	cancel() // Ensure all resources are cleaned up
}
