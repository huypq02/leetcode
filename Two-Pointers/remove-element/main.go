package main

import "fmt"

func removeElement(nums []int, val int) int {
	// Use two pointers:
	// 'low' for the next position to write a non-val element,
	// and 'fast' to scan through the array.
	low := 0
	// Iterate through the array with the 'fast' pointer.
	for fast := 0; fast < len(nums); fast++ {
		// If 'low' and 'fast' are not the same, copy the value at 'fast' to 'low'.
		if low != fast {
			nums[low] = nums[fast]
		}
		// If the current element is not equal to 'val', increment 'low'.
		if nums[low] != val {
			low++
		}
	}
	// 'low' now represents the number of elements not equal to 'val'.
	return low
}

func main() {
	// Example 1
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	expectedK1 := 2
	expectedNums1 := []int{2, 2}
	k1 := removeElement(nums1, val1)
	fmt.Println("Example 1:")
	fmt.Println("Returned k:", k1)
	fmt.Println("Expected k:", expectedK1)
	fmt.Println("First k elements:", nums1[:k1])
	fmt.Println("Expected first k elements:", expectedNums1)

	// Example 2
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	expectedK2 := 5
	expectedNums2 := []int{0, 1, 4, 0, 3} // Any order
	k2 := removeElement(nums2, val2)
	fmt.Println("Example 2:")
	fmt.Println("Returned k:", k2)
	fmt.Println("Expected k:", expectedK2)
	fmt.Println("First k elements:", nums2[:k2])
	fmt.Println("Expected first k elements (any order):", expectedNums2)
}
