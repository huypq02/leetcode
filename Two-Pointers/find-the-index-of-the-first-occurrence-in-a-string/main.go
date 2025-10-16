package main

func strStr(haystack string, needle string) int {
	// Check each possible starting index in haystack for a substring matching needle
	left := 0
	for right := len(needle); right <= len(haystack); right++ {

		// Check if the current substring matches needle
		if haystack[left:right] == needle {
			// If match found, return the starting index
			return left
		}
		left++
	}

	// If no match is found, return -1
	return -1
}

func main() {
	// Example 1
	haystack1 := "sadbutsad"
	needle1 := "sad"
	result1 := strStr(haystack1, needle1)
	println("Example 1:",
		"Input:", haystack1, needle1,
		"Output:", result1,
		"Expected:", 0)

	// Example 2
	haystack2 := "leetcode"
	needle2 := "leeto"
	result2 := strStr(haystack2, needle2)
	println("Example 2:",
		"Input:", haystack2, needle2,
		"Output:", result2,
		"Expected:", -1)

	// Example 3
	haystack3 := "a"
	needle3 := "a"
	result3 := strStr(haystack3, needle3)
	println("Example 3:",
		"Input:", haystack3, needle3,
		"Output:", result3,
		"Expected:", 0)

	// Example 4
	haystack4 := "abc"
	needle4 := "c"
	result4 := strStr(haystack4, needle4)
	println("Example 4:",
		"Input:", haystack4, needle4,
		"Output:", result4,
		"Expected:", 2)
}
