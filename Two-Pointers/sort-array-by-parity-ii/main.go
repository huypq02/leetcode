package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	// Initialize two pointers: evenPtr for even indices (0,2,4...), oddPtr for odd indices (1,3,5...)
	evenPtr, oddPtr := 0, 1

	// Continue until both pointers reach valid positions
	for evenPtr <= len(nums)-2 && oddPtr <= len(nums)-1 {
		// Skip correctly placed even numbers at even indices
		for evenPtr <= len(nums)-2 && nums[evenPtr]%2 == 0 {
			evenPtr += 2
		}

		// Skip correctly placed odd numbers at odd indices
		for oddPtr <= len(nums)-1 && nums[oddPtr]%2 != 0 {
			oddPtr += 2
		}

		// Only swap if both pointers are still within valid range
		if evenPtr <= len(nums)-2 && oddPtr <= len(nums)-1 {
			// Swap misplaced even number (at odd index) with misplaced odd number (at even index)
			nums[evenPtr], nums[oddPtr] = nums[oddPtr], nums[evenPtr]
		}
	}

	return nums
}

func main() {
	// Example 1: nums = [4,2,5,7]
	// Expected Output: [4,5,2,7] (or [4,7,2,5], [2,5,4,7], [2,7,4,5])
	nums1 := []int{4, 2, 5, 7}
	fmt.Printf("Input: %v\n", nums1)
	result1 := sortArrayByParityII(nums1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Expected: [4,5,2,7] (or [4,7,2,5], [2,5,4,7], [2,7,4,5])\n\n")

	// Example 2: nums = [2,3]
	// Expected Output: [2,3]
	nums2 := []int{2, 3}
	fmt.Printf("Input: %v\n", nums2)
	result2 := sortArrayByParityII(nums2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Expected: [2,3]\n\n")

	// Example 3: nums = [3,4]
	// Expected Output: [4,3]
	nums3 := []int{3, 4}
	fmt.Printf("Input: %v\n", nums3)
	result3 := sortArrayByParityII(nums3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Printf("Expected: [4,3]\n\n")
}
