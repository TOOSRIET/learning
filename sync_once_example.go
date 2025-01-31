package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	resource string
)

func initialize() {
	fmt.Println("Initializing resource...")
	resource = "Initialized"
}

func getResource() string {
	once.Do(initialize)
	return resource
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(getResource())
		}()
	}

	wg.Wait()
}
