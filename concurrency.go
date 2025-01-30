package main

import (
	"fmt"
	"sync"
	"time"
)

func expensiveDoubleFunction(j int) int {
	time.Sleep(1 * time.Second)
	return j * 2
}

func parrallelDouble(arr []int, workerNumber int) []int {
	n := len(arr)
	jobs := make(chan int, n)
	results := make(chan int, n)
	var wg sync.WaitGroup
	ans := []int{}

	// Start 3 workers
	for w := 1; w <= workerNumber; w++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func() {
			defer wg.Done()

			for j := range jobs {
				results <- expensiveDoubleFunction(j)
			}
		}()
	}

	for _, el := range arr {
		jobs <- el
	}

	close(jobs)

	go func() {
		wg.Wait() // Wait for all workers to finish
		close(results)
	}()

	for res := range results {
		ans = append(ans, res)
	}

	return ans
}

func main() {
	// Create test array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Measure sequential execution time
	startSeq := time.Now()
	seqResult := make([]int, len(arr))
	for i, v := range arr {
		seqResult[i] = expensiveDoubleFunction(v)
	}
	seqDuration := time.Since(startSeq)

	// Measure parallel execution time with 10 workers
	startPar := time.Now()
	parResult := parrallelDouble(arr, 10)
	parDuration := time.Since(startPar)

	// Print results and performance comparison
	fmt.Println("Sequential execution time:", seqDuration)
	fmt.Println(seqResult)
	fmt.Println("Parallel execution time:", parDuration)
	fmt.Println(parResult)
	fmt.Println("Performance improvement:", float64(seqDuration)/float64(parDuration), "x")

	// Output:
	// Sequential execution time: 10.009466042s
	// [2 4 6 8 10 12 14 16 18 20]
	// Parallel execution time: 1.0013415s
	// [8 6 10 12 16 20 4 2 18 14]
}
