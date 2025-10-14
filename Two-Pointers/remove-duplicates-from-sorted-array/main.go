package main

func removeDuplicates(nums []int) int {

	// Define low and fast pointers for removing duplicates
	// The low pointer keeps track of the position where the next unique element should be placed
	low := 0

	// The fast pointer iterates through the array to find unique elements
	for fast := 1; fast < len(nums); fast++ {
		// If we find a new unique element (different from the one at low position)
		if nums[low] != nums[fast] {
			// Move low pointer forward and place the unique element there
			low++
			nums[low] = nums[fast]
		}

	}

	// Return the count of unique elements
	return low + 1
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 2}
	expectedNums1 := []int{1, 2}
	k1 := removeDuplicates(nums1)
	println("Test Case 1:")
	println("Returned k:", k1)
	println("Expected k:", len(expectedNums1))
	println("First k elements:")
	for i := 0; i < k1; i++ {
		println(nums1[i])
	}

	// Example 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums2 := []int{0, 1, 2, 3, 4}
	k2 := removeDuplicates(nums2)
	println("\nTest Case 2:")
	println("Returned k:", k2)
	println("Expected k:", len(expectedNums2))
	println("First k elements:")
	for i := 0; i < k2; i++ {
		println(nums2[i])
	}
}

