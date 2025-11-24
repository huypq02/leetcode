package main

import "fmt"

func maximumCount(nums []int) int {
	// TODO: Implement solution
	return 0
}

func main() {
	// Test cases from README.md
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{-2, -1, -1, 1, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{-3, -2, -1, 0, 0, 1, 2},
			expected: 3,
		},
		{
			nums:     []int{5, 20, 66, 1314},
			expected: 4,
		},
	}

	for i, tc := range testCases {
		result := maximumCount(tc.nums)
		fmt.Printf("Test case %d: nums = %v\n", i+1, tc.nums)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
		fmt.Println()
	}
}
