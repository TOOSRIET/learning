package main

import (
	"fmt"
	"sync"
	"time"
)

// doubleElement doubles the given element and takes 1 second to complete.
func doubleElement(x int) int {
	time.Sleep(1 * time.Second)
	return x * 2
}

func add(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}

// sumOfDouble calculates the sum of double each element in the array using 5 workers.
func sumOfDouble(arr []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int, len(arr))
	sum := 0

	// Worker function to process elements from the channel
	worker := func(ch <-chan int) {
		defer wg.Done()
		for x := range ch {
			doubleX := doubleElement(x)
			mu.Lock()
			sum = add(sum, doubleX)
			mu.Unlock()
		}
	}

	// Start 5 workers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ch)
	}

	// Send elements to the channel
	for _, x := range arr {
		ch <- x
	}
	close(ch)

	// Wait for all workers to finish
	wg.Wait()

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	start := time.Now()
	sum := sumOfDouble(arr)
	elapsed := time.Since(start)

	fmt.Printf("Sum of double elements: %d\n", sum)
	fmt.Printf("Time taken: %s\n", elapsed)
}
