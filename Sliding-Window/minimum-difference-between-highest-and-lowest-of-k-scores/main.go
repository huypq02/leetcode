package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	// Sort the nums
	sort.Ints(nums)

	// Find the minimum difference the highest and the lowest of the k scores is minimized
	left, min := 0, math.MaxInt64
	for right := left + k - 1; right < len(nums); right++ {
		if min > nums[right]-nums[left] {
			min = nums[right] - nums[left]
		}
		left++
	}
	return min
}

func main() {
	// Example 1
	nums1 := []int{90}
	k1 := 1
	result1 := minimumDifference(nums1, k1)
	println("Example 1: Input: nums = [90], k = 1")
	println("Expected Output: 0")
	println("Actual Output:", result1)

	// Example 2
	nums2 := []int{9, 4, 1, 7}
	k2 := 2
	result2 := minimumDifference(nums2, k2)
	println("Example 2: Input: nums = [9,4,1,7], k = 2")
	println("Expected Output: 2")
	println("Actual Output:", result2)
}
