package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right, maxLength := 0, 1, 1
	indexSubstring := make(map[byte]int)
	indexSubstring[s[left]] = left
	// Iterate over a string
	for ; left < right && right < len(s); right++ {
		// Check the current substring exists or not
		if idx, ok := indexSubstring[s[right]]; ok {
			if left <= idx {
				left = idx + 1 // shift the left pointer to the index, which saves in the map before
			}
		}
		indexSubstring[s[right]] = right

		// Find the length of longest substring
		if maxLength < right-left+1 {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

func main() {
	// Test cases from README.md examples

	// Example 1: s = "abcabcbb", expected output: 3
	test1 := "abcabcbb"
	result1 := lengthOfLongestSubstring(test1)
	fmt.Printf("Example 1: Input: \"%s\", Output: %d, Expected: 3\n", test1, result1)

	// Example 2: s = "bbbbb", expected output: 1
	test2 := "bbbbb"
	result2 := lengthOfLongestSubstring(test2)
	fmt.Printf("Example 2: Input: \"%s\", Output: %d, Expected: 1\n", test2, result2)

	// Example 3: s = "pwwkew", expected output: 3
	test3 := "pwwkew"
	result3 := lengthOfLongestSubstring(test3)
	fmt.Printf("Example 3: Input: \"%s\", Output: %d, Expected: 3\n", test3, result3)

	// Additional edge cases

	// Edge case: empty string
	test4 := ""
	result4 := lengthOfLongestSubstring(test4)
	fmt.Printf("Edge case 1: Input: \"%s\", Output: %d, Expected: 0\n", test4, result4)

	// Edge case: single character
	test5 := "a"
	result5 := lengthOfLongestSubstring(test5)
	fmt.Printf("Edge case 2: Input: \"%s\", Output: %d, Expected: 1\n", test5, result5)

	// Test case: s = "au", expected output: 2
	test6 := "au"
	result6 := lengthOfLongestSubstring(test6)
	fmt.Printf("Test case 6: Input: \"%s\", Output: %d, Expected: 2\n", test6, result6)

	// Test case: s = "abba", expected output: 2
	test7 := "abba"
	result7 := lengthOfLongestSubstring(test7)
	fmt.Printf("Test case 7: Input: \"%s\", Output: %d, Expected: 2\n", test7, result7)
}
