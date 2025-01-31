package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	fmt.Println("Calling function A")
	functionA()
	fmt.Println("This will not be printed if functionA panics")
}

func functionA() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in functionA:", r)
		}
	}()

	fmt.Println("Calling function B")
	functionB()
	fmt.Println("This will not be printed if functionB panics")
}

func functionB() {
	fmt.Println("About to panic in functionB")
	panic("panic in functionB")
	fmt.Println("This will not be printed")
}
