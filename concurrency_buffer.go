package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ch := make(chan int, 5) buffer chan
	ch := make(chan int) // unbuffer chan
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		endTime := time.Now()
		fmt.Println("Time taken:", endTime.Sub(startTime))
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
