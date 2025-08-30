package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	// Clean data
	// lower
	s = strings.ToLower(s)
	// non-alphabet removal using the regular expressions technique
	re := regexp.MustCompile(`[^A-Za-z0-9]`)
	s = re.ReplaceAllString(s, "")

	// Create two pointers
	left, right := 0, len(s)-1
	// Loop until the two pointers meet each other
	// or the value of the two pointers is different
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	println(isPalindrome(s)) // Output: true

	s2 := "race a car"
	println(isPalindrome(s2)) // Output: false

}
