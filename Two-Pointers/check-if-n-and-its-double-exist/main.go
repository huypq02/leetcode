package main

import (
	"sort"
)

// binarySearch performs binary search to find if target exists in sorted array
func binarySearch(target int, array []int) bool {

	left, right := 0, len(array)-1

	for left <= right {
		mid := left + (right-left)/2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func checkIfExist(arr []int) bool {
	// Sort array to enable binary search
	sort.Ints(arr)
	n := len(arr)

	// For each element, search for its double in the remaining sorted array
	for i := 0; i < n; i++ {
		target := 2 * arr[i]
		// For negative even numbers, also check if half exists (since -10 is double of -5)
		if arr[i] < 0 && arr[i]%2 == 0 {
			target = arr[i] / 2
		}

		// Search in elements after current index to satisfy i != j
		if binarySearch(target, arr[i+1:]) {
			return true
		}
	}

	return false
}

func main() {
	// Example 1
	arr1 := []int{10, 2, 5, 3}
	// Expected output: true
	println("Test case 1: output =", checkIfExist(arr1), ", expected = true")

	// Example 2
	arr2 := []int{3, 1, 7, 11}
	// Expected output: false
	println("Test case 2: output =", checkIfExist(arr2), ", expected = false")

	// Example 3
	arr3 := []int{-10, 12, -20, -8, 15}
	// Expected output: true
	println("Test case 3: output =", checkIfExist(arr3), ", expected = true")
}
