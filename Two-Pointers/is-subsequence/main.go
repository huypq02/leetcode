package main

func isSubsequence(s string, t string) bool {
	// If s is empty, it's always a subsequence of t
	if len(s) == 0 {
		return true
	}

	// Use two pointers: left for s, i for t
	left, right := 0, len(s)-1

	// Iterate through t, try to match characters in s
	for i := 0; i < len(t); i++ {
		// If current character in t matches s[left], move left pointer
		if s[left] == t[i] {
			left++
		}

		// If all characters in s are matched, return true
		if left > right {
			return true
		}
	}

	// If not all characters in s are matched, return false
	return false
}

func main() {
	// Example 1
	s1 := "abc"
	t1 := "ahbgdc"
	println("Test case 1:", "Input:", s1, t1, "Expected Output: true")
	println("Actual Output:", isSubsequence(s1, t1))

	// Example 2
	s2 := "axc"
	t2 := "ahbgdc"
	println("Test case 2:", "Input:", s2, t2, "Expected Output: false")
	println("Actual Output:", isSubsequence(s2, t2))

	// Example 3: s is empty
	s3 := ""
	t3 := "ahbgdc"
	println("Test case 3:", "Input:", s3, t3, "Expected Output: true")
	println("Actual Output:", isSubsequence(s3, t3))

	// Example 4: s = "b", t = "c"
	s4 := "b"
	t4 := "c"
	println("Test case 4:", "Input:", s4, t4, "Expected Output: false")
	println("Actual Output:", isSubsequence(s4, t4))
}
