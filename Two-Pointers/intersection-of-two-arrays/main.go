package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {

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
			// If result is empty or last added element is different, append to result
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
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
	expected1 := []int{2}
	actual1 := intersection(nums1, nums2)
	fmt.Println("Example 1:")
	fmt.Println("Expected output:", expected1)
	fmt.Println("Actual output:", actual1)

	// Example 2
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	expected2 := []int{9, 4} // [4, 9] is also accepted
	actual2 := intersection(nums1, nums2)
	fmt.Println("Example 2:")
	fmt.Println("Expected output:", expected2, "or [4, 9]")
	fmt.Println("Actual output:", actual2)
}
