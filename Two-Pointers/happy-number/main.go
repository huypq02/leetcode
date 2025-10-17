package main

import (
	"strconv"
)

func isHappy(n int) bool {

	// Convert the input number to string for digit processing
	sumStr := strconv.Itoa(n)
	// Use a map to track numbers we've already seen to detect cycles
	seen := make(map[int]bool)

	for {
		// Convert the current string back to int
		sum, _ := strconv.Atoi(sumStr)

		// If the sum is 1, the number is happy
		if sum == 1 {
			return true
		}

		// If we've seen this sum before, we're in a cycle (not happy)
		if _, ok := seen[sum]; ok {
			return false
		}

		// Mark this sum as seen
		seen[sum] = true

		// Calculate the sum of the squares of the digits
		temp := 0
		for right := 0; right < len(sumStr); right++ {
			// Convert each digit character to int
			num, _ := strconv.Atoi(string(sumStr[right]))
			temp += num * num
		}

		// Prepare for next iteration with the new sum
		sumStr = strconv.Itoa(temp)
	}
}

func main() {
	// Example 1: Input: n = 19, Output: true
	n1 := 19
	expected1 := true
	result1 := isHappy(n1)
	println("Test case 1 - Input:", n1, "Expected:", expected1, "Got:", result1)

	// Example 2: Input: n = 2, Output: false
	n2 := 2
	expected2 := false
	result2 := isHappy(n2)
	println("Test case 2 - Input:", n2, "Expected:", expected2, "Got:", result2)

	// Additional test case: Input: n = 7, Expected: true
	n3 := 7
	expected3 := true
	result3 := isHappy(n3)
	println("Test case 3 - Input:", n3, "Expected:", expected3, "Got:", result3)
}
