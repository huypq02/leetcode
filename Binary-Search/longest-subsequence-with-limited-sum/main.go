package main

import (
	"fmt"
	"sort"
)

func binarySearch(prefix []int, target int) int {
	left, right := 0, len(prefix)-1

	// Find the rightmost position where prefix[mid] <= target
	for left <= right {
		mid := left + (right-left)/2

		if prefix[mid] <= target {
			// Current sum is valid, try to include more elements
			left = mid + 1
		} else {
			// Current sum exceeds target, try fewer elements
			right = mid - 1
		}
	}

	// left is now the count of elements we can include
	return left
}

func answerQueries(nums []int, queries []int) []int {
	// Sort nums in ascending order to maximize subsequence length
	// (greedy approach: take smallest elements first)
	sort.Ints(nums)

	n := len(nums)

	// Build prefix sum array for cumulative sums
	// Example: nums=[1,2,4,5] -> prefix=[1,3,7,12]
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// For each query, find maximum subsequence length using binary search
	answer := make([]int, len(queries))
	for i, q := range queries {
		// Binary search returns count of elements with sum <= q
		idx := binarySearch(prefix, q)
		answer[i] = idx
	}

	return answer
}

func main() {
	// Example 1
	nums1 := []int{4, 5, 2, 1}
	queries1 := []int{3, 10, 21}
	expected1 := []int{2, 3, 4}
	result1 := answerQueries(nums1, queries1)
	fmt.Println("Example 1:")
	fmt.Println("Input nums:", nums1)
	fmt.Println("Input queries:", queries1)
	fmt.Println("Expected:", expected1)
	fmt.Println("Result:", result1)

	// Example 2
	nums2 := []int{2, 3, 4, 5}
	queries2 := []int{1}
	expected2 := []int{0}
	result2 := answerQueries(nums2, queries2)
	fmt.Println("\nExample 2:")
	fmt.Println("Input nums:", nums2)
	fmt.Println("Input queries:", queries2)
	fmt.Println("Expected:", expected2)
	fmt.Println("Result:", result2)
}
