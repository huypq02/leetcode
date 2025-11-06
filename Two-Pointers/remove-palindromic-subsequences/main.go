package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
	return 0
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
