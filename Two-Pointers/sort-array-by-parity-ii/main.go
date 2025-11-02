package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	
	return nums
}

func main() {
	// Example 1: nums = [4,2,5,7]
	// Expected Output: [4,5,2,7] (or [4,7,2,5], [2,5,4,7], [2,7,4,5])
	nums1 := []int{4, 2, 5, 7}
	fmt.Printf("Input: %v\n", nums1)
	result1 := sortArrayByParityII(nums1)
	fmt.Printf("Output: %v\n\n", result1)

	// Example 2: nums = [2,3]
	// Expected Output: [2,3]
	nums2 := []int{2, 3}
	fmt.Printf("Input: %v\n", nums2)
	result2 := sortArrayByParityII(nums2)
	fmt.Printf("Output: %v\n\n", result2)
}
