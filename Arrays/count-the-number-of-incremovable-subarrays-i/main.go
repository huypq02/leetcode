package main

import (
	"fmt"
)

// buildPrefixOk precomputes whether nums[0:i+1] is strictly increasing
// prefixOk[i] = true means nums[0] < nums[1] < ... < nums[i]
// This avoids redundant checks in the main loop
func buildPrefixOk(nums []int) []bool {
	n := len(nums)
	prefixOk := make([]bool, n)

	// Base case: single element always strictly increasing
	prefixOk[0] = true

	// Dynamic programming: prefixOk[i] depends on prefixOk[i-1]
	for i := 1; i < n; i++ {
		if prefixOk[i-1] && nums[i-1] < nums[i] {
			prefixOk[i] = true
		} else {
			prefixOk[i] = false
		}
	}
	return prefixOk
}

// buildSuffixOk precomputes whether nums[i:n] is strictly increasing
// suffixOk[i] = true means nums[i] < nums[i+1] < ... < nums[n-1]
// This avoids redundant checks in the main loop
func buildSuffixOk(nums []int) []bool {
	n := len(nums)
	suffixOk := make([]bool, n)

	// Base case: last element
	suffixOk[n-1] = true

	// Dynamic programming: suffixOk[i] depends on suffixOk[i+1]
	for i := n - 2; i >= 0; i-- {
		if suffixOk[i+1] && nums[i] < nums[i+1] {
			suffixOk[i] = true
		} else {
			suffixOk[i] = false
		}
	}
	return suffixOk
}

func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	count := 0

	// Precompute prefix and suffix validity to avoid redundant checks
	prefixOk := buildPrefixOk(nums)
	suffixOk := buildSuffixOk(nums)

	// Try every possible subarray [i, j] to remove
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// After removing [i, j], we have:
			// - prefix: nums[0:i] (empty if i == 0)
			// - suffix: nums[j+1:n] (empty if j == n-1)

			// Check if prefix nums[0:i] is strictly increasing
			prefixOk := (i == 0) || prefixOk[i-1]

			// Check if suffix nums[j+1:n] is strictly increasing
			suffixOk := (j == n-1) || suffixOk[j+1]

			// Check if last element of prefix < first element of suffix
			// This ensures concatenating prefix + suffix is strictly increasing
			connectionOk := true
			if i > 0 && j < n-1 {
				connectionOk = nums[i-1] < nums[j+1]
			}

			// If all three conditions are met, this subarray is incremovable
			if prefixOk && suffixOk && connectionOk {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3, 4}
	expected1 := 10
	result1 := incremovableSubarrayCount(nums1)
	fmt.Printf("Test 1: nums = %v\n", nums1)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected1, result1, result1 == expected1)

	// Test case 2
	nums2 := []int{6, 5, 7, 8}
	expected2 := 7
	result2 := incremovableSubarrayCount(nums2)
	fmt.Printf("Test 2: nums = %v\n", nums2)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected2, result2, result2 == expected2)

	// Test case 3
	nums3 := []int{8, 7, 6, 6}
	expected3 := 3
	result3 := incremovableSubarrayCount(nums3)
	fmt.Printf("Test 3: nums = %v\n", nums3)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected3, result3, result3 == expected3)

	// Test case 4
	nums4 := []int{2, 1, 6}
	expected4 := 5
	result4 := incremovableSubarrayCount(nums4)
	fmt.Printf("Test 4: nums = %v\n", nums4)
	fmt.Printf("Expected: %d, Got: %d, Pass: %v\n\n", expected4, result4, result4 == expected4)
}
