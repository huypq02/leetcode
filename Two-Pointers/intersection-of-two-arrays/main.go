package main

func intersection(nums1 []int, nums2 []int) []int {
	return []int{}
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	// Expected output: [2]
	_ = intersection(nums1, nums2)

	// Example 2
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	// Expected output: [9, 4] or [4, 9]
	_ = intersection(nums1, nums2)
}
