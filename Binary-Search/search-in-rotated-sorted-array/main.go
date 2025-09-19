package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// On the left part of the array sorted in ascending order
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // On the right part of the array sorted in ascending order
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {

	// Example 1
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	result1 := search(nums1, target1)
	println("Example 1 Output:", result1) // Expected: 4

	// Example 2
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	result2 := search(nums2, target2)
	println("Example 2 Output:", result2) // Expected: -1

	// Example 3
	nums3 := []int{1}
	target3 := 0
	result3 := search(nums3, target3)
	println("Example 3 Output:", result3) // Expected: -1

	// Additional Testcase
	nums4 := []int{1}
	target4 := 1
	result4 := search(nums4, target4)
	println("Additional Testcase Output:", result4) // Expected: 0

	// Additional Testcase
	nums5 := []int{1, 3}
	target5 := 3
	result5 := search(nums5, target5)
	println("Additional Testcase Output:", result5) // Expected: 1

	// Additional Testcase
	nums6 := []int{1, 3}
	target6 := 1
	result6 := search(nums6, target6)
	println("Additional Testcase Output:", result6) // Expected: 0

	// Additional Testcase
	nums7 := []int{3, 5, 1}
	target7 := 3
	result7 := search(nums7, target7)
	println("Additional Testcase Output:", result7) // Expected: 0

}
