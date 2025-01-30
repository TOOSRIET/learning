package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	participantA := make(chan int)
	participantB := make(chan int)

	go func() {
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(10) + 1
		fmt.Println("A", randomNum)
		time.Sleep(time.Duration(randomNum * 1000))
		participantA <- 1
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(10) + 1
		fmt.Println("B", randomNum)
		time.Sleep(time.Duration(randomNum * 1000))
		participantB <- 1
	}()

	select {
	case <-participantA:
		println("A won")
	case <-participantB:
		println("B won")
	case <-time.After(15 * time.Second):
		println("Timeout")
	}
}
