package main

import "fmt"

func diStringMatch(s string) []int {
	// Two-pointer approach: maintain min and max values
	// For 'I': use the smallest available number (ensures perm[i] < perm[i+1])
	// For 'D': use the largest available number (ensures perm[i] > perm[i+1])

	n := len(s)
	perm := make([]int, n+1) // Result array of size n+1
	min, max := 0, n         // Initialize two pointers at both ends of range [0, n]

	// Process each character in the string
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'I':
			// Increasing: assign the minimum value and increment min pointer
			perm[i] = min
			min++
		case 'D':
			// Decreasing: assign the maximum value and decrement max pointer
			perm[i] = max
			max--
		}
	}

	// At the end, min and max converge to the same value
	// Assign this remaining value to the last position
	perm[n] = min

	return perm
}

func main() {
	// Example 1
	s1 := "IDID"
	expected1 := []int{0, 4, 1, 3, 2}
	result1 := diStringMatch(s1)
	fmt.Println("Input:", s1)
	fmt.Println("Expected Output:", expected1)
	fmt.Println("Actual Output:", result1)

	// Example 2
	s2 := "III"
	expected2 := []int{0, 1, 2, 3}
	result2 := diStringMatch(s2)
	fmt.Println("Input:", s2)
	fmt.Println("Expected Output:", expected2)
	fmt.Println("Actual Output:", result2)

	// Example 3
	s3 := "DDI"
	expected3 := []int{3, 2, 0, 1}
	result3 := diStringMatch(s3)
	fmt.Println("Input:", s3)
	fmt.Println("Expected Output:", expected3)
	fmt.Println("Actual Output:", result3)
}
