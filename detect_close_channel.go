package main

import (
	"fmt"
	"time"
)

func isClosedChannel(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int)
	close(ch)

	go func() {
		for el := range ch {
			fmt.Println(el)
		}
	}()

	if isClosedChannel(ch) {
		fmt.Println("Channel is closed")
		return
	}

	ch <- 1

	time.Sleep(2 * time.Second)
}
