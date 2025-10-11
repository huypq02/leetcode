package main

func longestNiceSubstring(s string) string {
	maxLen, start := 0, 0
	// Try all possible substrings
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substr := s[i:j]
			// Update result if a longer nice substring is found
			if isNice(substr) && len(substr) > maxLen {
				maxLen = len(substr)
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

// isNice checks if a string is "nice" according to the problem definition.
func isNice(s string) bool {
	// Track presence of each uppercase and lowercase letter
	upper := make([]bool, 26)
	lower := make([]bool, 26)

	// Mark which letters are present in uppercase/lowercase
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			upper[c-'A'] = true
		} else if c >= 'a' && c <= 'z' {
			lower[c-'a'] = true
		}
	}

	// For each letter, if present, must appear in both cases
	for i := 0; i < 26; i++ {
		if upper[i] != lower[i] {
			return false
		}
	}

	return true
}

func main() {
	// Example 1
	s1 := "YazaAay"
	expected1 := "aAa"
	println("Input:", s1)
	println("Expected Output:", expected1)
	println("Actual Output:", longestNiceSubstring(s1))
	println()

	// Example 2
	s2 := "Bb"
	expected2 := "Bb"
	println("Input:", s2)
	println("Expected Output:", expected2)
	println("Actual Output:", longestNiceSubstring(s2))
	println()

	// Example 3
	s3 := "c"
	expected3 := ""
	println("Input:", s3)
	println("Expected Output:", expected3)
	println("Actual Output:", longestNiceSubstring(s3))
	println()
}
