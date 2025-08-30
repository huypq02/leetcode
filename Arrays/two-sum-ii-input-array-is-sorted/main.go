package main

import (
	"fmt"
)


func twoSum(numbers []int, target int) []int {
	// Create 2 pointers left and right
	left, right := 0, len(numbers)-1

	// Retrieving each value in the array
	for left < right {
		sum := numbers[left] + numbers[right] // Sum two numbers at the pointers
		switch {
		case sum < target:
			left++
		case sum > target:
			right--
		default: // The sum is equal to the target
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

func main() {
	// Input: nums = [2,7,11,15], target = 9
	// Output: [1,2]
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1))

	// Input: nums = [2,3,4], target = 6
	// Output: [1,3]
	nums2 := []int{2, 3, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))

	// Input: nums = [-1,0], target = -1
	// Output: [1,2]
	nums3 := []int{-1, 0}
	target3 := -1
	fmt.Println(twoSum(nums3, target3))
}
