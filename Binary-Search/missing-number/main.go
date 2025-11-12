package main

import (
	"sort"
)

func missingNumber(nums []int) int {
	// Sort the array in ascending order to ensure indices match values if no number is missing before them
	sort.Ints(nums)

	// Initialize binary search boundaries
	left, right := 0, len(nums)-1
	// Initialize answer to len(nums) in case the missing number is n (the last number in the range)
	answer := len(nums)

	// Perform binary search
	for left <= right {
		mid := left + (right-left)/2
		// If the value at index mid equals mid, the missing number is to the right
		if nums[mid] == mid {
			left = mid + 1
		} else {
			// If not, the missing number is at mid or to the left
			answer = mid
			right = mid - 1
		}
	}

	// Return the missing number
	return answer
}

func main() {
	// Example 1
	nums1 := []int{3, 0, 1}
	// Expected output: 2
	println("Test case 1: nums = [3,0,1], output =", missingNumber(nums1), ", expected = 2")

	// Example 2
	nums2 := []int{0, 1}
	// Expected output: 2
	println("Test case 2: nums = [0,1], output =", missingNumber(nums2), ", expected = 2")

	// Example 3
	nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	// Expected output: 8
	println("Test case 3: nums = [9,6,4,2,3,5,7,0,1], output =", missingNumber(nums3), ", expected = 8")

	// Example 4
	nums4 := []int{1, 2}
	// Expected output: 0
	println("Test case 4: nums = [1,2], output =", missingNumber(nums4), ", expected = 0")

	// Example 5
	nums5 := []int{1, 2, 3}
	// Expected output: 0
	println("Test case 5: nums = [1,2,3], output =", missingNumber(nums5), ", expected = 0")

}
