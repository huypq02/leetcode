package main

import (
	"sort"
)

// binarySearch returns the first index in sorted nums where nums[idx] >= x.
// If all elements are less than x, returns len(nums).
// This is a lower_bound implementation (first position >= x).
func binarySearch(nums []int, x int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if x <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// specialArray finds the unique integer x such that exactly x elements in nums are >= x.
// If no such x exists, returns -1.
// Approach:
//  1. Sort nums in non-decreasing order.
//  2. Use binary search on possible x (from 0 to n) to find the unique x.
//     For each x, use lower_bound (binarySearch) to count how many elements are >= x.
//     If count == x, return x. If count < x, search smaller x. If count > x, search larger x.
func specialArray(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	left, right := 0, n
	for left <= right {
		x := left + (right-left)/2
		// Find the first index where nums[idx] >= x
		idx := binarySearch(nums, x)
		count := n - idx // Number of elements >= x

		if count == x {
			return x
		} else if count < x {
			right = x - 1 // Too few, try smaller x
		} else {
			left = x + 1 // Too many, try larger x
		}
	}

	return -1
}

func main() {
	// Example 1
	nums1 := []int{3, 5}
	println("Output:", specialArray(nums1), "(Expected: 2)")

	// Example 2
	nums2 := []int{0, 0}
	println("Output:", specialArray(nums2), "(Expected: -1)")

	// Example 3
	nums3 := []int{0, 4, 3, 0, 4}
	println("Output:", specialArray(nums3), "(Expected: 3)")

	// Additional Testcase
	nums4 := []int{3, 9, 7, 8, 3, 8, 6, 6}
	println("Output:", specialArray(nums4), "(Expected: 6)")
}
