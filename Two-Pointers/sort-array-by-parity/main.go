package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	// Maintain slow as the next slot for an even value and sweep with fast.
	slow, fast := 0, 0
	for fast < len(nums) {
		// When we see an even number, swap it into the slow position and advance slow.
		if nums[fast]%2 == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		// Always progress fast so every element is visited exactly once.
		fast++
	}

	return nums
}

func main() {
	// Example 1
	nums1 := []int{3, 1, 2, 4}
	expected1 := []int{2, 4, 3, 1} // or any valid permutation
	actual1 := sortArrayByParity(nums1)
	fmt.Println("Example 1:")
	fmt.Println("Expected output:", expected1)
	fmt.Println("Actual output:", actual1)

	// Example 2
	nums2 := []int{0}
	expected2 := []int{0}
	actual2 := sortArrayByParity(nums2)
	fmt.Println("Example 2:")
	fmt.Println("Expected output:", expected2)
	fmt.Println("Actual output:", actual2)

	// Example 3
	nums3 := []int{0, 1, 2}
	expected3 := []int{0, 2, 1}
	actual3 := sortArrayByParity(nums3)
	fmt.Println("Example 3:")
	fmt.Println("Expected output:", expected3)
	fmt.Println("Actual output:", actual3)
}
