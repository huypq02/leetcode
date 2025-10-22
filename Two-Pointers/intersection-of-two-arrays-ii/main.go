package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	// Sort both arrays in ascending order to enable two-pointer technique
	sort.Ints(nums1)
	sort.Ints(nums2)

	// Initialize two pointers for each array
	i, j := 0, 0

	n1, n2 := len(nums1), len(nums2)
	result := []int{}
	// Traverse both arrays with two pointers
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			// Append to result
			result = append(result, nums1[i])
			// Move both pointers forward
			i++
			j++
		} else if nums1[i] < nums2[j] {
			// Move pointer i forward if nums1[i] is smaller
			i++
		} else { // nums1[i] > nums2[j]
			// Move pointer j forward if nums2[j] is smaller
			j++
		}
	}

	return result
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result1 := intersect(nums1, nums2)
	fmt.Println("Example 1 Output:", result1, "Expected: [2 2]")

	// Example 2
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	result2 := intersect(nums1, nums2)
	fmt.Println("Example 2 Output:", result2, "Expected: [4 9] or [9 4]")
}
