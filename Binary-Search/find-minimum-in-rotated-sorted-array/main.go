package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	// Loop until the left and right pointers meet
	for left < right {
		mid := (right-left)/2 + left
		last := nums[right] // the last element of the array
		switch {
		case nums[mid] < last:
			right = mid // the minimum element is on the left of the array
		case nums[mid] > last:
			left = mid + 1 // the minimum element is on the right of the array
		}
	}
	return nums[left]
}

func main() {
	// Test cases
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{3, 4, 5, 1, 2}, 1},
		{"Example 2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"Example 3", []int{11, 13, 15, 17}, 11},
		{"Example 4", []int{3, 1, 2}, 1},
		{"Example 5", []int{2, 3, 4, 5, 1}, 1},
		{"Example 6", []int{2, 1}, 1},
		{"Example 7", []int{5, 1, 2, 3, 4}, 1},
	}

	for _, tc := range testCases {
		println(tc.name)
		println("Input:", fmt.Sprint(tc.nums))
		output := findMin(tc.nums)
		if output != tc.expected {
			println("FALSE:", output, "\b, expected output:", tc.expected)
		} else {
			println("TRUE:", tc.expected)
		}
		println("")
	}
}
