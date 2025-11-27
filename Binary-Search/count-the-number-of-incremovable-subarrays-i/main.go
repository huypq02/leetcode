package main

import "fmt"

func incremovableSubarrayCount(nums []int) int {
	// TODO: implement the solution to solve the exercise here
	return 0
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
}
