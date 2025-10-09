package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	// left points to the start of the window; sum tracks the sum of the current window
	left, sum := 0, 0
	// maxSum initialized to minimum int to handle arrays with negative values
	maxSum := math.MinInt64

	// Iterate right pointer to expand window
	for right := 0; right < len(nums); right++ {
		// Add the new element to the current window sum
		sum += nums[right]

		// If window size exceeds k, shrink from left and update sum
		if right-left+1 > k {
			sum -= nums[left]
			left++
		}

		// Only update maxSum when window size exactly k
		if right-left+1 == k && maxSum < sum {
			maxSum = sum
		}
	}

	// Return the maximum average found; uses float64 division for precision
	return float64(maxSum) / float64(k)
}

func main() {
	// Example 1
	nums1 := []int{1, 12, -5, -6, 50, 3}
	k1 := 4
	result1 := findMaxAverage(nums1, k1)
	println("Example 1 Output:", result1, "Expected: 12.75")

	// Example 2
	nums2 := []int{5}
	k2 := 1
	result2 := findMaxAverage(nums2, k2)
	println("Example 2 Output:", result2, "Expected: 5.0")

	// Example 3
	nums3 := []int{-6662, 5432, -8558, -8935, 8731, -3083, 4115, 9931, -4006, -3284, -3024, 1714, -2825, -2374, -2750, -959, 6516, 9356, 8040, -2169, -9490, -3068, 6299, 7823, -9767, 5751, -7897, 6680, -1293, -3486, -6785, 6337, -9158, -4183, 6240, -2846, -2588, -5458, -9576, -1501, -908, -5477, 7596, -8863, -4088, 7922, 8231, -4928, 7636, -3994, -243, -1327, 8425, -3468, -4218, -364, 4257, 5690, 1035, 6217, 8880, 4127, -6299, -1831, 2854, -4498, -6983, -677, 2216, -1938, 3348, 4099, 3591, 9076, 942, 4571, -4200, 7271, -6920, -1886, 662, 7844, 3658, -6562, -2106, -296, -3280, 8909, -8352, -9413, 3513, 1352, -8825}
	k3 := 90
	result3 := findMaxAverage(nums3, k3)
	println("Example 3 Output:", result3, "Expected: 37.25556")

	// Example 4
	nums4 := []int{4, 0, 4, 3, 3}
	k4 := 5
	result4 := findMaxAverage(nums4, k4)
	println("Example 4 Output:", result4, "Expected: 2.8")

	// Example 5
	nums5 := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	k5 := 93
	result5 := findMaxAverage(nums5, k5)
	println("Example 5 Output:", result5, "Expected: -594.58065")
}
