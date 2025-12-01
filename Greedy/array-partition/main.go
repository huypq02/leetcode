package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	// Greedy approach: Sort the array in ascending order
	// To maximize the sum of minimums, pair adjacent elements after sorting
	// This ensures we don't waste larger numbers by pairing them with much smaller ones
	sort.Ints(nums)

	// Sum up every other element starting from index 0
	// After sorting, the element at even indices will be the minimum of each pair
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}

func main() {
	// Test case 1
	nums1 := []int{1, 4, 3, 2}
	result1 := arrayPairSum(nums1)
	println("Test case 1: Input:", nums1, "Output:", result1)

	// Test case 2
	nums2 := []int{6, 2, 6, 5, 1, 2}
	result2 := arrayPairSum(nums2)
	println("Test case 2: Input:", nums2, "Output:", result2)
}
