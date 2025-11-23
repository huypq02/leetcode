package main

func answerQueries(nums []int, queries []int) []int {
	return []int{}
}

func main() {
	// Example 1
	nums1 := []int{4, 5, 2, 1}
	queries1 := []int{3, 10, 21}
	expected1 := []int{2, 3, 4}
	result1 := answerQueries(nums1, queries1)
	println("Example 1:")
	println("Input nums:", nums1)
	println("Input queries:", queries1)
	println("Expected:", expected1)
	println("Result:", result1)

	// Example 2
	nums2 := []int{2, 3, 4, 5}
	queries2 := []int{1}
	expected2 := []int{0}
	result2 := answerQueries(nums2, queries2)
	println("\nExample 2:")
	println("Input nums:", nums2)
	println("Input queries:", queries2)
	println("Expected:", expected2)
	println("Result:", result2)
}
