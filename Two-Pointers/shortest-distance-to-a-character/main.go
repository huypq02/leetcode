package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	// Initialize result array to store distances
	result := make([]int, len(s))
	// Use nearest to track the closest index of character c found so far
	nearest := math.MaxInt64
	left, right := 0, len(s)-1

	// First pass: left to right
	// For each index, update nearest if c is found, and store distance to nearest c
	for left <= right {
		if s[left] == c {
			nearest = left
		}
		// Distance from current index to nearest c on the left
		result[left] = nearest - left
		if result[left] < 0 {
			result[left] = -result[left]
		}
		left++
	}

	// Second pass: right to left
	// For each index, update nearest if c is found, and store the minimum distance to c
	nearest = math.MaxInt64
	left = 0
	for left <= right {
		if s[right] == c {
			nearest = right
		}
		// Update result with the minimum distance to c from either direction
		result[right] = min(nearest-right, result[right])
		right--
	}

	return result
}

func main() {
	// Example 1
	s1 := "loveleetcode"
	c1 := byte('e')
	expected1 := []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}
	result1 := shortestToChar(s1, c1)
	fmt.Println("Testcase 1:")
	fmt.Println("Input:", s1, string(c1))
	fmt.Println("Expected:", expected1)
	fmt.Println("Output:", result1)

	// Example 2
	s2 := "aaab"
	c2 := byte('b')
	expected2 := []int{3, 2, 1, 0}
	result2 := shortestToChar(s2, c2)
	fmt.Println("Testcase 2:")
	fmt.Println("Input:", s2, string(c2))
	fmt.Println("Expected:", expected2)
	fmt.Println("Output:", result2)

	// Example 3
	s3 := "aaba"
	c3 := byte('b')
	expected3 := []int{2, 1, 0, 1}
	result3 := shortestToChar(s3, c3)
	fmt.Println("Testcase 3:")
	fmt.Println("Input:", s3, string(c3))
	fmt.Println("Expected:", expected3)
	fmt.Println("Output:", result3)
}
