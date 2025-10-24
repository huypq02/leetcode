package main

import "sort"

func findContentChildren(g []int, s []int) int {
	// Handle edge cases: if no children or no cookies, return 0
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	// Sort both arrays to enable greedy two-pointer approach
	// Sorting allows us to match smallest greed with smallest cookie
	sort.Ints(g) // greed factors in ascending order
	sort.Ints(s) // cookie sizes in ascending order

	// Initialize two pointers: i for children, j for cookies
	i, j := 0, 0

	// Two-pointer technique: traverse both arrays simultaneously
	for i < len(g) && j < len(s) {
		// If current cookie satisfies current child's greed
		if g[i] <= s[j] {
			// Child i is satisfied, move to next child
			i++
			// Use this cookie, move to next cookie
			j++
		} else {
			// Current cookie is too small for child i
			// Try the next cookie without moving to next child
			// This is the key insight: we skip small cookies, not children
			j++
		}
	}

	// Return count of satisfied children (value of pointer i)
	// i represents how many children have been satisfied (0 to len(g))
	return i
}

func main() {
	// Example 1
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1}
	result1 := findContentChildren(g1, s1)
	println("Example 1 Output:", result1, "Expected: 1")

	// Example 2
	g2 := []int{1, 2}
	s2 := []int{1, 2, 3}
	result2 := findContentChildren(g2, s2)
	println("Example 2 Output:", result2, "Expected: 2")
}
