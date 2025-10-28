package main

func validPalindrome(s string) bool {
	// Helper function to check if s[left:right] can be a palindrome with at most one deletion
	var canBePalindrome func(s string, left, right, deleted int) bool
	canBePalindrome = func(s string, left, right, deleted int) bool {
		for left < right {
			// If characters at both ends do not match
			if s[left] != s[right] {
				deleted++ // Increment deletion count
				// If more than one deletion is needed, return false
				if deleted > 1 {
					return false
				}
				// Try skipping either the left or right character and check recursively
				return canBePalindrome(s, left, right-1, deleted) || canBePalindrome(s, left+1, right, deleted)
			}
			// Move both pointers towards the center
			left++
			right--
		}
		// If loop completes, it's a palindrome (with at most one deletion)
		return true
	}

	// Start checking from both ends with zero deletions
	return canBePalindrome(s, 0, len(s)-1, 0)
}

func main() {
	// Example 1
	s1 := "aba"
	println("Input:", s1, "Output:", validPalindrome(s1), "Expected: true")

	// Example 2
	s2 := "abca"
	println("Input:", s2, "Output:", validPalindrome(s2), "Expected: true")

	// Example 3
	s3 := "abc"
	println("Input:", s3, "Output:", validPalindrome(s3), "Expected: false")

	// Example 4
	s4 := "deeee"
	println("Input:", s4, "Output:", validPalindrome(s4), "Expected: true")

	// Example 5
	s5 := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	println("Input:", s5, "Output:", validPalindrome(s5), "Expected: true")

	// Example 6
	s6 := "eceec"
	println("Input:", s6, "Output:", validPalindrome(s6), "Expected: true")
}
