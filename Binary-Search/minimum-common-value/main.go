package main

func getCommon(nums1 []int, nums2 []int) int {
	// TODO: Implement my solution here to solve the exercise
	return 0
}

func main() {
	// Example 1
	nums11 := []int{1, 2, 3}
	nums21 := []int{2, 4}
	result1 := getCommon(nums11, nums21)
	println("Example 1:", result1, "Expected: 2")

	// Example 2
	nums12 := []int{1, 2, 3, 6}
	nums22 := []int{2, 3, 4, 5}
	result2 := getCommon(nums12, nums22)
	println("Example 2:", result2, "Expected: 2")

	// Test case: No common element
	nums13 := []int{1, 3, 5}
	nums23 := []int{2, 4, 6}
	result3 := getCommon(nums13, nums23)
	println("No common element:", result3, "Expected: -1")

	// Test case: Single element arrays with common value
	nums14 := []int{5}
	nums24 := []int{5}
	result4 := getCommon(nums14, nums24)
	println("Single element:", result4, "Expected: 5")

	// Test case: Large arrays with common element at the end
	nums15 := []int{1, 2, 3, 4, 100}
	nums25 := []int{50, 60, 70, 80, 100}
	result5 := getCommon(nums15, nums25)
	println("Common at end:", result5, "Expected: 100")
}
