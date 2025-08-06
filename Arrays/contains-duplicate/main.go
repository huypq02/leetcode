package main

import "fmt"

// containsDuplicate checks if there are any duplicate elements in the given slice of integers.
func containsDuplicate(nums []int) bool {
	count := make(map[int]int, len(nums))
	for _, v := range nums {
		count[v]++
		if count[v] > 1 {
			return true
		}
	}
	return false

}

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1)) // Output: true
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) // Output: false
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3)) // Output: true
	nums4 := []int{1, 2, 3, 4, 5}
	fmt.Println(containsDuplicate(nums4))
}
