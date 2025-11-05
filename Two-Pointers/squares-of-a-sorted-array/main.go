package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	// Use two pointers to compare absolute values from both ends of the array.
	// Place the larger square at the end of the result array, moving inward.
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	for left <= right {
		// Compare squares of left and right elements.
		if nums[left]*nums[left] < nums[right]*nums[right] {
			// Place the larger square at the current position from the end.
			result[right-left] = nums[right] * nums[right]
			right--
		} else {
			result[right-left] = nums[left] * nums[left]
			left++
		}
	}

	// The result array is filled in non-decreasing order.
	return result
}

func main() {
	// Example 1
	nums1 := []int{-4, -1, 0, 3, 10}
	expected1 := []int{0, 1, 9, 16, 100}
	fmt.Println("Input:", nums1)
	fmt.Println("Expected Output:", expected1)
	fmt.Println("Actual Output:", sortedSquares(nums1))

	// Example 2
	nums2 := []int{-7, -3, 2, 3, 11}
	expected2 := []int{4, 9, 9, 49, 121}
	fmt.Println("Input:", nums2)
	fmt.Println("Expected Output:", expected2)
	fmt.Println("Actual Output:", sortedSquares(nums2))
}
