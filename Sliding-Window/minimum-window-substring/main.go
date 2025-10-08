package main

func minWindow(s string, t string) string {
	// Return "" if t cannot be a substring of s
	if len(s) < len(t) {
		return ""
	}

	// Map to count required characters in t
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// Map to count current window characters in s
	window := make(map[byte]int)
	left, right := 0, 0
	valid, start, minLen := 0, 0, len(s)+1

	for ; right < len(s); right++ {
		// Expand window by including s[right]
		if count, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == count {
				valid++
			}
		}

		// Shrink window from the left while valid
		for valid == len(need) && left <= right {
			// Update minimum window if smaller found
			if minLen > len(s[left:right+1]) {
				minLen = len(s[left : right+1])
				start = left
			}

			// Remove s[left] from window
			if count, ok := need[s[left]]; ok {
				// Decrement the count of s[left] in the current window
				window[s[left]]--
				// If the count falls below what's needed, decrement valid
				if window[s[left]] < count {
					valid--
				}
			}
			left++
		}

	}

	// Return result or "" if not found
	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}

func main() {
	// Example 1
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	println("Example 1:")
	println("Input: s = \"" + s1 + "\", t = \"" + t1 + "\"")
	println("Expected Output: BANC")
	println("Actual Output:", minWindow(s1, t1))
	println()

	// Example 2
	s2 := "a"
	t2 := "a"
	println("Example 2:")
	println("Input: s = \"" + s2 + "\", t = \"" + t2 + "\"")
	println("Expected Output: a")
	println("Actual Output:", minWindow(s2, t2))
	println()

	// Example 3
	s3 := "a"
	t3 := "aa"
	println("Example 3:")
	println("Input: s = \"" + s3 + "\", t = \"" + t3 + "\"")
	println("Expected Output: ")
	println("Actual Output:", minWindow(s3, t3))
}
