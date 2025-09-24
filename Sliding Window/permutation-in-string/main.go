package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	// Use array to store lowercase English letters only
	counterS1, counterS2 := [26]int{}, [26]int{}

	for i := 0; i < len(s1); i++ {
		// Update frequency of current character of s1
		counterS1[s1[i]-'a']++
	}

	left, right := 0, len(s1)-1
	for right < len(s2) {

		// Update frequency of current character of s2
		for j := left; j <= right; j++ {
			counterS2[s2[j]-'a']++
		}

		// Compare two arrays if they equals return true
		if counterS2 == counterS1 {
			return true
		}
		// Reset the counter of s2
		counterS2 = [26]int{}
		// Shift window size
		left++
		right++
	}

	return false
}

func main() {

	// Example 1
	s1 := "ab"
	s2 := "eidbaooo"
	result := checkInclusion(s1, s2)
	fmt.Printf("Example 1:\nInput: s1 = \"%s\", s2 = \"%s\"\nOutput: %t\nExpected: true\n\n", s1, s2, result)

	// Example 2
	s1 = "ab"
	s2 = "eidboaoo"
	result = checkInclusion(s1, s2)
	fmt.Printf("Example 2:\nInput: s1 = \"%s\", s2 = \"%s\"\nOutput: %t\nExpected: false\n\n", s1, s2, result)

}
