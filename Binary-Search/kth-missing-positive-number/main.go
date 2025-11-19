package main

import (
	"fmt"
	"reflect"
)

func findKthPositive(arr []int, k int) int {
	// Complete sequence: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...
	// Array present:        2, 3, 4,       7,           11

	// Missing numbers:   1,          5, 6,    8, 9, 10,    12, 13, ...
	//                    ↑           ↑  ↑     ↑  ↑   ↑
	//                   1st         2nd 3rd  4th 5th 6th ...

	// Binary search to find the smallest index where the count of missing numbers is >= k
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		// Number of missing positive integers before arr[mid]
		// For arr[mid], the count of missing numbers is arr[mid] - (mid + 1)
		missingCount := arr[mid] - (mid + 1)
		if missingCount < k {
			// If missing numbers so far is less than k, move right
			left = mid + 1
		} else {
			// Otherwise, move left
			right = mid - 1
		}
	}

	// After the loop, right is the last index where missingCount < k
	// The answer is k + right + 1
	return k + right + 1
}

func main() {
	// Example 1
	arr1 := []int{2, 3, 4, 7, 11}
	k1 := 5
	expected1 := 9
	result1 := findKthPositive(arr1, k1)
	fmt.Printf("Test 1 - Input: arr = %v, k = %d\n", arr1, k1)
	fmt.Printf("Expected: %d, Got: %d\n", expected1, result1)
	fmt.Println("Pass:", reflect.DeepEqual(result1, expected1))

	// Example 2
	arr2 := []int{1, 2, 3, 4}
	k2 := 2
	expected2 := 6
	result2 := findKthPositive(arr2, k2)
	fmt.Printf("Test 2 - Input: arr = %v, k = %d\n", arr2, k2)
	fmt.Printf("Expected: %d, Got: %d\n", expected2, result2)
	fmt.Println("Pass:", reflect.DeepEqual(result2, expected2))
}
