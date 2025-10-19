package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	// Two-pointer approach:
	// 'fast' scans every element in the array.
	// 'slow' marks the position to place the next non-zero element.
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// Swap the current non-zero element at 'fast' with the element at 'slow'.
			// This moves non-zero elements forward and zeros to the end.
			temp := nums[slow]
			nums[slow] = nums[fast]
			nums[fast] = temp
			// Increment 'slow' to the next position for a non-zero element.
			slow++
		}
		// Increment 'fast' to continue scanning the array.
		fast++
	}
}

func main() {
	// Example 1
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println("Example 1 result:", nums1, "Expected: [1 3 12 0 0]")

	// Example 2
	nums2 := []int{0}
	moveZeroes(nums2)
	fmt.Println("Example 2 result:", nums2, "Expected: [0]")
}
