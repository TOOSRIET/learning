package main

import (
	"fmt"
	"sync"
	"time"
)

func doubleElement(x int) int {
	time.Sleep(1 * time.Second)
	return x * 2
}

func tripleElement(x int) int {
	time.Sleep(1 * time.Second)
	return x * 3
}

func add(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}

func processElements(arr []int, transformer func(int) int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int, len(arr))
	sum := 0

	worker := func(ch <-chan int) {
		defer wg.Done()
		for x := range ch {
			transformed := transformer(x)
			mu.Lock()
			sum = add(sum, transformed)
			mu.Unlock()
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ch)
	}

	for _, x := range arr {
		ch <- x
	}
	close(ch)

	wg.Wait()
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup
	var doubleSum, tripleSum int

	start := time.Now()

	// Start double sum calculation
	wg.Add(1)
	go func() {
		defer wg.Done()
		doubleSum = processElements(arr, doubleElement)
	}()

	// Start triple sum calculation
	wg.Add(1)
	go func() {
		defer wg.Done()
		tripleSum = processElements(arr, tripleElement)
	}()

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Sum of double elements: %d\n", doubleSum)
	fmt.Printf("Sum of triple elements: %d\n", tripleSum)
	fmt.Printf("Time taken: %s\n", elapsed)
}
