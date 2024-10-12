package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge the two arrays and find the median
	merged := append(nums1, nums2...)
	// sort the two arrays in a single array
	sort.Ints(merged)
	// find the median
	fmt.Println(merged)
	length := len(merged) / 2

	return float64(merged[length]) + float64(merged[length+1])/2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
