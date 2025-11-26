package main

import (
	"fmt"
	"sort"
)

// lowerBound finds the first index where nums[index] >= search
// Returns the count of elements strictly less than search
func lowerBound(nums []int, search int) int {
	left, right := 0, len(nums)

	// Binary search to find the insertion point
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < search {
			left = mid + 1 // Move right if mid element is too small
		} else {
			right = mid // Move left if mid element is >= search
		}
	}
	// left now points to the first element >= search, which equals count of smaller elements
	return left
}

func countPairs(nums []int, target int) int {
	// Sort the array to enable binary search
	sort.Ints(nums)

	count := 0
	// For each element, find how many elements after it can form valid pairs
	for i := 0; i < len(nums)-1; i++ {
		// Calculate what value j needs to be less than: nums[i] + nums[j] < target
		// Therefore: nums[j] < target - nums[i]
		// Use binary search to count elements in nums[i+1:] that are <  target - nums[i]
		count += lowerBound(nums[i+1:], target-nums[i])
	}

	return count
}

func main() {
	// Test case 1
	nums1 := []int{-1, 1, 2, 3, 1}
	target1 := 2
	expected1 := 3
	result1 := countPairs(nums1, target1)
	fmt.Printf("Test 1: nums=%v, target=%d\n", nums1, target1)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected1, result1, result1 == expected1)

	// Test case 2
	nums2 := []int{-6, 2, 5, -2, -7, -1, 3}
	target2 := -2
	expected2 := 10
	result2 := countPairs(nums2, target2)
	fmt.Printf("Test 2: nums=%v, target=%d\n", nums2, target2)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected2, result2, result2 == expected2)
}
