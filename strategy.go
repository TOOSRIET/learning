package main

import (
	"fmt"
)

type SortStrategy interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (b BubbleSort) Sort(arr []int) []int {
	fmt.Println("Sorting using Bubble Sort")
	return arr
}

type QuickSort struct{}

func (q QuickSort) Sort(arr []int) []int {
	fmt.Println("Sorting using Quick Sort")
	return arr
}

type SortContext struct {
	strategy SortStrategy
}

func (s *SortContext) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func main() {
	fmt.Println("Hello, World!")

	context := &SortContext{}
	arr := []int{5, 2, 8, 1, 9}
	context.SetStrategy(BubbleSort{})
	sortedArr := context.strategy.Sort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
