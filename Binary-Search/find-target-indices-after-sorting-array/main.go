package main

import (
	"fmt"
	"sort"
)

func lowerBound(nums []int, target int) int {
	// Finds the first index where nums[index] >= target
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// Returns the lower bound index
	return left
}

func upperBound(nums []int, target int) int {
	// Finds the first index where nums[index] > target
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// Returns the upper bound index
	return left
}

func targetIndices(nums []int, target int) []int {
	// Sort the array to prepare for binary search
	sort.Ints(nums)

	// Find the range of indices where nums[index] == target
	lowerBound := lowerBound(nums, target) // first index >= target
	upperBound := upperBound(nums, target) // first index > target

	// Build the result slice with all indices where nums[index] == target
	targetIdx := make([]int, upperBound-lowerBound)
	for i := range targetIdx {
		targetIdx[i] = lowerBound + i
	}

	return targetIdx
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 5, 2, 3}
	target1 := 2
	result1 := targetIndices(nums1, target1)
	fmt.Println("Example 1 Input before sorting:", []int{1, 2, 5, 2, 3}, "Target:", target1, "Output:", result1, "Expected: [1 2]")

	// Example 2
	nums2 := []int{1, 2, 5, 2, 3}
	target2 := 3
	result2 := targetIndices(nums2, target2)
	fmt.Println("Example 2 Input before sorting:", []int{1, 2, 5, 2, 3}, "Target:", target2, "Output:", result2, "Expected: [3]")

	// Example 3
	nums3 := []int{1, 2, 5, 2, 3}
	target3 := 5
	result3 := targetIndices(nums3, target3)
	fmt.Println("Example 3 Input before sorting:", []int{1, 2, 5, 2, 3}, "Target:", target3, "Output:", result3, "Expected: [4]")

	// Example 4
	nums4 := []int{100, 1, 100}
	target4 := 100
	result4 := targetIndices(nums4, target4)
	fmt.Println("Example 4 Input before sorting:", []int{100, 1, 100}, "Target:", target4, "Output:", result4, "Expected: [1 2]")
}
