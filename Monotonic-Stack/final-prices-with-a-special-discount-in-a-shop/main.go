package main

import "fmt"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Append a new value to the slice (dynamic array)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1] // Truncate the slice by reducing its length by one element
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1] // Top is the last element of the slice
}

func finalPrices(prices []int) []int {
	n := len(prices)
	answer := prices
	discountIdx := Stack{}

	for i := 0; i < n; i++ {
		for len(discountIdx) > 0 && prices[discountIdx.Top()] >= prices[i] {
			prevIdx := discountIdx.Top()
			discountIdx.Pop()
			answer[prevIdx] = prices[prevIdx] - prices[i]
		}
		discountIdx.Push(i)
	}

	return answer
}

func main() {
	// Test cases
	test1 := []int{8, 4, 6, 2, 3}
	test2 := []int{1, 2, 3, 4, 5}
	test3 := []int{10, 1, 1, 6}

	fmt.Printf("Test 1: %v -> %v\n", test1, finalPrices(test1))
	fmt.Printf("Test 2: %v -> %v\n", test2, finalPrices(test2))
	fmt.Printf("Test 3: %v -> %v\n", test3, finalPrices(test3))
}
