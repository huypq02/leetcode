package main

import "fmt"

func findDuplicate(nums []int) int {
	// Define search range: we're looking for numbers in range [1, n]
	left, right := 1, len(nums)

	// Binary search on the VALUE range [1, n], not on array indices
	for left < right {
		mid := (left + right) / 2

		// Count how many numbers in nums are ≤ mid
		// If count > mid, then by pigeonhole principle,
		// there must be a duplicate in the range [1, mid]
		if countNum(nums, mid) > mid {
			right = mid // Duplicate is in range [left, mid]
		} else {
			// If count ≤ mid, the duplicate must be in range [mid+1, right]
			left = mid + 1
		}
	}

	// When left == right, we've found the duplicate number
	return left // or return right because left is equal to right
}

func countNum(nums []int, length int) int {
	count := 0
	for _, n := range nums {
		if n <= length {
			count++
		}
	}
	return count
}

func main() {
	// Test cases from README.md examples

	// Example 1
	nums1 := []int{1, 3, 4, 2, 2}
	result1 := findDuplicate(nums1)
	fmt.Printf("Example 1: nums = %v\n", nums1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 2\n\n")

	// Example 2
	nums2 := []int{3, 1, 3, 4, 2}
	result2 := findDuplicate(nums2)
	fmt.Printf("Example 2: nums = %v\n", nums2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 3\n\n")

	// Example 3
	nums3 := []int{3, 3, 3, 3, 3}
	result3 := findDuplicate(nums3)
	fmt.Printf("Example 3: nums = %v\n", nums3)
	fmt.Printf("Output: %d\n", result3)
	fmt.Printf("Expected: 3\n\n")
}
