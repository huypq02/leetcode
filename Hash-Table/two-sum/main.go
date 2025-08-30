package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		if _, ok := hash[v]; ok {
			return []int{hash[v], i}
		}
		hash[target-v] = i
	}
	return []int{-1, -1}
}

func main() {
	// Input: nums = [2,7,11,15], target = 9
	nums, target := []int{2, 7, 11, 15}, 9
	// Output: [0,1]
	fmt.Println(twoSum(nums, target))

	// Input: nums = [3,2,4], target = 6
	nums2, target2 := []int{3, 2, 4}, 6
	// Output: [1,2]
	fmt.Println(twoSum(nums2, target2))

	// Input: nums = [3,3], target = 6
	nums3, target3 := []int{3, 3}, 6
	// Output: [0,1]
	fmt.Println(twoSum(nums3, target3))

}
