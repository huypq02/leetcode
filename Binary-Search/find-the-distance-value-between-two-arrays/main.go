package main

import "sort"

func findTargetInRange(arr2 []int, lower, upper int) bool {
	left, right := 0, len(arr2)-1
	for left <= right {
		mid := left + (right-left)/2
		// If arr2[mid] is within [lower, upper], we found a value in range
		if arr2[mid] >= lower {
			if arr2[mid] <= upper {
				return true
			}
			// arr2[mid] > upper, search left half
			right = mid - 1
		} else {
			// arr2[mid] < lower, search right half
			left = mid + 1
		}
	}
	// No value found in range
	return false
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// Sort arr2 to enable binary search
	sort.Ints(arr2)

	count := 0
	for _, num := range arr1 {
		// Check if any value in arr2 is within [num-d, num+d]
		if findTargetInRange(arr2, num-d, num+d) {
			// If we find any value in arr2 within [num-d, num+d],
			// it means |arr1[i] - arr2[j]| <= d for some j, so arr1[i] does NOT satisfy the distance condition.
			// Therefore, we skip counting this num.
			continue
		}
		// If not found, increment the count
		count++
	}
	return count
}

func main() {
	// Example 1
	arr1_1 := []int{4, 5, 8}
	arr2_1 := []int{10, 9, 1, 8}
	d1 := 2
	println("Test Case 1 Output:", findTheDistanceValue(arr1_1, arr2_1, d1), "| Expected: 2")

	// Example 2
	arr1_2 := []int{1, 4, 2, 3}
	arr2_2 := []int{-4, -3, 6, 10, 20, 30}
	d2 := 3
	println("Test Case 2 Output:", findTheDistanceValue(arr1_2, arr2_2, d2), "| Expected: 2")

	// Example 3
	arr1_3 := []int{2, 1, 100, 3}
	arr2_3 := []int{-5, -2, 10, -3, 7}
	d3 := 6
	println("Test Case 3 Output:", findTheDistanceValue(arr1_3, arr2_3, d3), "| Expected: 1")
}
