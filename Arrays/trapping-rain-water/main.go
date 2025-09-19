package main

import (
	"fmt"
)

func trap(height []int) int {
	// For each index i, trapped water is calculated as:
	// water_at_i = min(left_max[i], right_max[i]) - height[i]
	// left_max[i]: max height from 0 to i
	// right_max[i]: max height from i to n-1
	// Example for height = [0,1,0,2,1,0,1,3,2,1,2,1]:
	// i=2: left_max=1, right_max=3, height=0, water=1
	// i=5: left_max=2, right_max=3, height=0, water=2
	// ...

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	water := 0
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	// Using Prefix Maximum to find the left max at each position
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	// Using Suffix Maximum to find the right max at each position
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// Calculate trapped water at each position
	for i := 1; i < n-1; i++ {
		waterAtIdx := min(leftMax[i], rightMax[i]) - height[i]
		if waterAtIdx < 0 {
			continue
		}
		water += waterAtIdx
	}
	return water
}

func main() {
	testCases := []struct {
		height   []int
		expected int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, expected: 9},
	}

	for i, tc := range testCases {
		result := trap(tc.height)
		if result != tc.expected {
			fmt.Printf("Test case %d failed: input=%v, expected=%d, got=%d\n", i+1, tc.height, tc.expected, result)
		} else {
			fmt.Printf("Test case %d success: input=%v, got=%d\n", i+1, tc.height, result)
		}
	}
}
