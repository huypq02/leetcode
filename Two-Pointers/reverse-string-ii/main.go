package main

func reverseStr(s string, k int) string {

	// Convert the input string to a slice of runes to handle Unicode characters safely
	runes := []rune(s)

	// Iterate over the string in windows of size 2*k
	for start := 0; start < len(runes); start += 2 * k {
		// For each window, set left and right pointers to reverse the first k characters
		left := start
		right := min(start+k-1, len(runes)-1) // Ensure we don't go out of bounds

		// Reverse the characters between left and right
		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	// Convert the runes slice back to a string and return
	return string(runes)
}

func main() {
	// Example 1
	s1 := "abcdefg"
	k1 := 2
	expected1 := "bacdfeg"
	result1 := reverseStr(s1, k1)
	println("Input:", s1, "k:", k1)
	println("Expected Output:", expected1)
	println("Actual Output:", result1)
	println()

	// Example 2
	s2 := "abcd"
	k2 := 2
	expected2 := "bacd"
	result2 := reverseStr(s2, k2)
	println("Input:", s2, "k:", k2)
	println("Expected Output:", expected2)
	println("Actual Output:", result2)
	println()

	// Example 3
	s3 := "abcdef"
	k3 := 3
	expected3 := "cbadef"
	result3 := reverseStr(s3, k3)
	println("Input:", s3, "k:", k3)
	println("Expected Output:", expected3)
	println("Actual Output:", result3)
	println()

	// Example 4
	s4 := "abcdefg"
	k4 := 8
	expected4 := "gfedcba"
	result4 := reverseStr(s4, k4)
	println("Input:", s4, "k:", k4)
	println("Expected Output:", expected4)
	println("Actual Output:", result4)
	println()

	// Example 5
	s5 := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	k5 := 39
	expected5 := "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi"
	result5 := reverseStr(s5, k5)
	println("Input:", s5, "k:", k5)
	println("Expected Output:", expected5)
	println("Actual Output:", result5)
	println()
}
