package main

import (
	"fmt"
)

func arrangeCoins(n int) int {

	// We use binary search to find the largest k such that k*(k+1)/2 <= n
	// This represents the maximum number of complete rows that can be formed
	left, right := 1, n
	result := 0
	for left <= right {
		// k is the current guess for the number of complete rows
		k := left + (right-left)/2
		// Calculate the total coins needed to form k complete rows
		total := k * (k + 1) / 2

		if total <= n {
			// If we can form k rows, try for more rows
			result = k
			left = k + 1
		} else {
			// If not enough coins, try fewer rows
			right = k - 1
		}
	}

	// result holds the maximum number of complete rows
	return result
}

func main() {
	testCases := []struct {
		n        int
		expected int
	}{
		{5, 2},
		{8, 3},
		{2, 1},
	}

	for i, tc := range testCases {
		result := arrangeCoins(tc.n)
		fmt.Printf("Test Case %d: n = %d\n", i+1, tc.n)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASSED")
		} else {
			fmt.Println("FAILED")
		}
		fmt.Println()
	}
}
