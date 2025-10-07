package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m, n = len(nums1), len(nums2)
	var total = m + n

	var findKth func(nums1 []int, nums2 []int, k int) int
	findKth = func(nums1 []int, nums2 []int, k int) int {
		// Always ensure nums1 is the shorter array for optimization
		if len(nums1) > len(nums2) {
			return findKth(nums2, nums1, k)
		}
		// If nums1 is empty, return the kth element from nums2
		if len(nums1) == 0 {
			return nums2[k-1]
		}

		if k == 1 {
			return min(nums1[0], nums2[0])
		}

		// Using binary search for finding kth element efficiently
		i := min(k/2, len(nums1))
		j := k - i
		if nums1[i-1] < nums2[j-1] {
			return findKth(nums1[i:], nums2, k-i)
		}
		return findKth(nums1, nums2[j:], k-j)
	}

	// For odd total length, the median is the middle element
	if total%2 == 1 {
		return float64(findKth(nums1, nums2, total/2+1))
	}
	// For even total length, median is the average of the two middle elements
	left := findKth(nums1, nums2, total/2)
	right := findKth(nums1, nums2, total/2+1)
	return float64(left+right) / 2.0
}

func main() {
	// Example 1:
	nums1 := []int{1, 3}
	nums2 := []int{2}
	// Expected output: 2.00000
	// Explanation: merged array = [1,2,3] and median is 2.
	println("Example 1:")
	println("Input: nums1 = [1,3], nums2 = [2]")
	println("Expected output: 2.00000")
	println("Output:", findMedianSortedArrays(nums1, nums2))

	// Example 2:
	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	// Expected output: 2.50000
	// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
	println("Example 2:")
	println("Input: nums1 = [1,2], nums2 = [3,4]")
	println("Expected output: 2.50000")
	println("Output:", findMedianSortedArrays(nums3, nums4))

}
