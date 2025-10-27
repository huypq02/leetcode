package main

import (
	"unicode"
)

func reverseRange(runes []rune, start, end int) {
	for start < end {
		// Swap characters at start and end
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	start := 0
	// Loop through each character
	for end := 0; end < len(runes); end++ {
		// If end of word, reverse it
		if end == len(runes)-1 || unicode.IsSpace(runes[end+1]) {
			reverseRange(runes, start, end)
			start = end + 2 // Move to next word
		}
	}

	return string(runes)
}

func main() {
	// Example 1
	input1 := "Let's take LeetCode contest"
	expected1 := "s'teL ekat edoCteeL tsetnoc"
	result1 := reverseWords(input1)
	println("Input:", input1)
	println("Expected Output:", expected1)
	println("Actual Output:", result1)
	println()

	// Example 2
	input2 := "Mr Ding"
	expected2 := "rM gniD"
	result2 := reverseWords(input2)
	println("Input:", input2)
	println("Expected Output:", expected2)
	println("Actual Output:", result2)
}
