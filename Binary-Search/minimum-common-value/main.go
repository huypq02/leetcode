package main

// binarySearch performs binary search on a sorted array to find the target value
// Returns the target if found, otherwise returns -1
func binarySearch(nums []int, target int) int {
	// Initialize left and right pointers for binary search
	left, right := 0, len(nums)
	for left < right {
		// Calculate mid to avoid overflow
		mid := left + (right-left)/2
		if nums[mid] == target {
			// Target found, return it
			return nums[mid]
		} else if nums[mid] < target {
			// Target is in the right half
			left = mid + 1
		} else {
			// Target is in the left half
			right = mid
		}
	}
	// Target not found
	return -1
}

// getCommon finds the minimum common value between two sorted arrays
// Approach: Iterate through nums1 and use binary search to check if each element exists in nums2
func getCommon(nums1 []int, nums2 []int) int {
	// Iterate through each element in nums1 (in ascending order due to sorted array)
	for _, v := range nums1 {
		// Use binary search to check if current element exists in nums2
		if binarySearch(nums2, v) != -1 {
			return v
		}
	}
	// No common element found
	return -1
}

func main() {
	// Example 1
	nums11 := []int{1, 2, 3}
	nums21 := []int{2, 4}
	result1 := getCommon(nums11, nums21)
	println("Example 1:", result1, "Expected: 2")

	// Example 2
	nums12 := []int{1, 2, 3, 6}
	nums22 := []int{2, 3, 4, 5}
	result2 := getCommon(nums12, nums22)
	println("Example 2:", result2, "Expected: 2")

	// Test case: No common element
	nums13 := []int{1, 3, 5}
	nums23 := []int{2, 4, 6}
	result3 := getCommon(nums13, nums23)
	println("No common element:", result3, "Expected: -1")

	// Test case: Single element arrays with common value
	nums14 := []int{5}
	nums24 := []int{5}
	result4 := getCommon(nums14, nums24)
	println("Single element:", result4, "Expected: 5")

	// Test case: Large arrays with common element at the end
	nums15 := []int{1, 2, 3, 4, 100}
	nums25 := []int{50, 60, 70, 80, 100}
	result5 := getCommon(nums15, nums25)
	println("Common at end:", result5, "Expected: 100")
}
