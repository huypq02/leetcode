package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums) // Sort to get largest sides at the end

	// Iterate from the largest side, checking consecutive triplets
	for i := len(nums) - 2; i > 0; i-- {
		// Triangle inequality: sum of two smaller sides > largest side
		if nums[i]+nums[i-1] > nums[i+1] {
			return nums[i] + nums[i-1] + nums[i+1]
		}
	}

	return 0 // No valid triangle found
}

func main() {
	// Test case 1
	nums1 := []int{2, 1, 2}
	fmt.Println("Input:", nums1)
	// Expected Output: 5
	fmt.Println("Expected Output: 5")
	fmt.Println("Actual Output:", largestPerimeter(nums1))

	// Test case 2
	nums2 := []int{1, 2, 1, 10}
	fmt.Println("Input:", nums2)
	// Expected Output: 0
	fmt.Println("Expected Output: 0")
	fmt.Println("Actual Output:", largestPerimeter(nums2))
}
