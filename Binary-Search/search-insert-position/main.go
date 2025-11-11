package main

import (
	"fmt"
	"strings"
)

func searchInsert(nums []int, target int) int {
	// Use binary search to find the target or the position to insert
	left, right := 0, len(nums)-1
	// left: start index, right: end index
	for left <= right {
		mid := left + (right-left)/2 // Find the middle index
		if nums[mid] < target {
			// Target is in the right half
			left = mid + 1
		} else if nums[mid] > target {
			// Target is in the left half
			right = mid - 1
		} else {
			// Target found at mid
			return mid
		}
	}
	// If not found, left is the correct insert position
	return left
}

func main() {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3}, 2, 1},
	}

	for i, tc := range testcases {
		numsStr := make([]string, len(tc.nums))
		for j, v := range tc.nums {
			numsStr[j] = fmt.Sprintf("%d", v)
		}
		fmt.Printf("Test case %d: nums = [%s], target = %d, output = %d, expected = %d\n",
			i+1, strings.Join(numsStr, ","), tc.target, searchInsert(tc.nums, tc.target), tc.expected)
	}
}
