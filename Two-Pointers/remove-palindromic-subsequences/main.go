package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
	// Initialize step counter (this represents the current attempt number)
	stepCounter := 1

	// Define a nested function to check palindrome possibility with deletions
	// Parameters: s = string, left = left pointer, right = right pointer, deleted = deletion count
	var canBePalindrome func(s string, left, right, deleted int) int
	canBePalindrome = func(s string, left, right, deleted int) int {
		// Use two pointers approach: check from both ends moving inward
		for left < right {
			// When characters at current positions don't match
			if s[left] != s[right] {
				deleted++ // Count this as a required deletion

				// If we need more than one deletion, return the count
				// (This logic seems to limit deletions to 1, which doesn't match the problem)
				if deleted > 1 {
					return deleted
				}
			}
			// Move pointers toward center for next comparison
			left++
			right--
		}
		// All characters matched (with allowed deletions) - return deletion count
		return deleted
	}

	// Start palindrome check from string boundaries with initial step count
	return canBePalindrome(s, 0, len(s)-1, stepCounter)
}

func main() {
	// Example 1
	s1 := "ababa"
	fmt.Println("Input:", s1)
	fmt.Println("Expected Output: 1")
	fmt.Println("Actual Output:", removePalindromeSub(s1))
	fmt.Println()

	// Example 2
	s2 := "abb"
	fmt.Println("Input:", s2)
	fmt.Println("Expected Output: 2")
	fmt.Println("Actual Output:", removePalindromeSub(s2))
	fmt.Println()

	// Example 3
	s3 := "baabb"
	fmt.Println("Input:", s3)
	fmt.Println("Expected Output: 2")
	fmt.Println("Actual Output:", removePalindromeSub(s3))
	fmt.Println()
}
