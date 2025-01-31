package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu  sync.RWMutex // Read-Write Mutex
	val int
}

func expensiveSum(a int, b int) int {
	time.Sleep(1 * time.Second)
	return a + b
}

// Increment updates the counter (needs Write Lock)
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val = expensiveSum(c.val, 1)
}

// Value reads the counter (needs Read Lock)
func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val
}

func main() {
	counter := SafeCounter{}

	var wg sync.WaitGroup

	// Start multiple goroutines for reading and writing
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Value:", counter.Value())
		}()
	}

	wg.Wait()
}
