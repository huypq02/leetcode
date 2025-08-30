package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Sort the array
	sort.Ints(nums)
	// Validate zero values
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	// Define the length of the array
	n := len(nums)
	// Define a list of the result
	resultList := [][]int{}
	// Retrieving the array and adding the proper result to the list of the result
	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Define two pointers
		left, right := i+1, n-1
		target := -nums[i]
		for left < right {
			sum := nums[left] + nums[right] // Sum two numbers at the pointers
			switch {
			case sum < target:
				left++
			case sum > target:
				right--
			default: // The sum is equal to the target
				resultList = append(resultList, []int{nums[i], nums[left], nums[right]})
				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			}
		}
	}

	return resultList
}

func main() {
	// First test case
	// Input: nums = [-1,0,1,2,-1,-4]
	// Output: [[-1,-1,2],[-1,0,1]]
	nums := []int{-1, 0, 1, 2, -1, -4}
	// Call the function
	fmt.Println("this is the result: ", threeSum(nums))

	// Second
	// Input [0,0,0,0]
	// Output [0,0,0,0]
	nums2 := []int{0, 0, 0, 0}
	// Call the function
	fmt.Println("this is the result: ", threeSum(nums2))

	// Third
	// Input: [2,-3,0,-2,-5,-5,-4,1,2,-2,2,0,2,-4,5,5,-10]
	// Output: [[-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]]
	nums3 := []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}
	// Call the function
	fmt.Println("this is the result: ", threeSum(nums3))

	// Input: [0,-4,-1,-4,-2,-3,2]
	// Output: [[-2,0,2]]
	nums4 := []int{0, -4, -1, -4, -2, -3, 2}
	// Call the function
	fmt.Println("this is the result: ", threeSum(nums4))
}
