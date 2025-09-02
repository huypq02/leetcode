package main

func search(nums []int, target int) int {
	// Define the left and right pointers
	left, right := 0, len(nums)-1

	// Loop until the left and right pointers meet
	for left <= right {
		mid := (right-left)/2 + left
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default: // if existing the 'target' in the 'nums' array
			return mid
		}
	}

	return -1
}

func main() {
	// Input: nums = [-1,0,3,5,9,12], target = 9
	nums1, target1 := []int{-1, 0, 3, 5, 9, 12}, 9
	// Output: 4
	println(search(nums1, target1))

	// 	Input: nums = [-1,0,3,5,9,12], target = 2
	nums2, target2 := []int{-1, 0, 3, 5, 9, 12}, 2
	// Output: -1
	println(search(nums2, target2))

}
