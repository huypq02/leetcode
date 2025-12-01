package main

import (
	"fmt"
)

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	used := make(map[int]bool) // Track which baskets are already used
	count := 0                 // Count of unplaced fruits
	for _, f := range fruits {
		placed := false
		// Try to place the current fruit in the leftmost available basket
		for i, b := range baskets {
			if !used[i] && f <= b {
				used[i] = true // Mark basket as used
				placed = true
				break // Stop after placing in the first suitable basket
			}
		}
		if !placed {
			count++ // Fruit could not be placed
		}
	}
	return count
}

func main() {
	// Example 1
	fruits1 := []int{4, 2, 5}
	baskets1 := []int{3, 5, 4}
	expected1 := 1
	result1 := numOfUnplacedFruits(fruits1, baskets1)
	fmt.Printf("Example 1:\n")
	fmt.Printf("Input: fruits = %v, baskets = %v\n", fruits1, baskets1)
	fmt.Printf("Output: %d, Expected: %d\n", result1, expected1)
	if result1 == expected1 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	fmt.Println()

	// Example 2
	fruits2 := []int{3, 6, 1}
	baskets2 := []int{6, 4, 7}
	expected2 := 0
	result2 := numOfUnplacedFruits(fruits2, baskets2)
	fmt.Printf("Example 2:\n")
	fmt.Printf("Input: fruits = %v, baskets = %v\n", fruits2, baskets2)
	fmt.Printf("Output: %d, Expected: %d\n", result2, expected2)
	if result2 == expected2 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	fmt.Println()

	// Example 3
	fruits3 := []int{7, 5}
	baskets3 := []int{8, 5}
	expected3 := 0
	result3 := numOfUnplacedFruits(fruits3, baskets3)
	fmt.Printf("Example 3:\n")
	fmt.Printf("Input: fruits = %v, baskets = %v\n", fruits3, baskets3)
	fmt.Printf("Output: %d, Expected: %d\n", result3, expected3)
	if result3 == expected3 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
