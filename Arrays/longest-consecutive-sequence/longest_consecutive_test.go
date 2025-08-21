package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	// List of testcases
	testcases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"the Longest Consecutive Sequence of [100,4,200,1,3,2]",
			[]int{100, 4, 200, 1, 3, 2}, 4},
		{"the Longest Consecutive Sequence of [0,3,7,2,5,8,4,6,0,1]",
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"the Longest Consecutive Sequence of [1,0,1,2]",
			[]int{1, 0, 1, 2}, 3},
		{"sequence is nil or empty",
			[]int{}, 0},
		{"the Longest Consecutive Sequence of [9,1,4,7,3,-1,0,5,8,-1,6]",
			[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
		{"the Longest Consecutive Sequence of [9,1,-3,2,4,8,3,-1,6,-2,-4,7]",
			[]int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}, 4},
	}

	// Retrieving all of the testcase
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestConsecutive(tc.nums)
			if result != tc.expected {
				t.Errorf("The Longest Consecutive Sequence of %v = %v; want %v", tc.nums, result, tc.expected)
			}
		})
	}
}
