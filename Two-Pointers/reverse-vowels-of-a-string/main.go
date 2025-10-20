package main

import (
	"fmt"
)

func swap(runes []rune, left, right int) {
	// Swap the characters at positions left and right in the rune slice
	runes[left], runes[right] = runes[right], runes[left]
}

func reverseVowels(s string) string {
	// Map to quickly check if a character is a vowel (case-insensitive)
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Convert the string to a slice of runes to handle Unicode characters safely
	runes := []rune(s)
	left, right := 0, len(runes)-1

	// Two-pointer approach: move left and right pointers towards each other
	for left < right {
		// If both pointers are at vowels, swap them
		if vowels[runes[left]] && vowels[runes[right]] {
			swap(runes, left, right)
		} else if vowels[runes[left]] {
			// If only left is vowel, move right pointer leftwards
			right--
			continue
		} else if vowels[runes[right]] {
			// If only right is vowel, move left pointer rightwards
			left++
			continue
		}
		// If neither is vowel, move both pointers
		left++
		right--
	}

	// Convert the rune slice back to string and return
	return string(runes)
}

func main() {
	// Testcase 1
	s1 := "IceCreAm"
	expected1 := "AceCreIm"
	fmt.Printf("Input: %q\nExpected Output: %q\nActual Output: %q\n\n", s1, expected1, reverseVowels(s1))

	// Testcase 2
	s2 := "leetcode"
	expected2 := "leotcede"
	fmt.Printf("Input: %q\nExpected Output: %q\nActual Output: %q\n\n", s2, expected2, reverseVowels(s2))

	// Testcase 3
	s3 := "Sore was I ere I saw Eros."
	expected3 := "SorE was I ere I saw eros."
	fmt.Printf("Input: %q\nExpected Output: %q\nActual Output: %q\n\n", s3, expected3, reverseVowels(s3))
}
