package main

func countBinarySubstrings(s string) int {
	// prevRun: length of previous run of consecutive characters
	// currRun: length of current run of consecutive characters
	// result: total count of valid substrings
	prevRun, currRun := 0, 1
	result := 0
	// Iterate through the string, grouping consecutive 0's or 1's
	for right := 0; right < len(s)-1; right++ {
		if s[right] == s[right+1] {
			// If current and next char are the same, extend current run
			currRun++
			continue
		}
		// When the character changes, count valid substrings
		// The number of valid substrings is the minimum of previous and current run lengths
		result += min(prevRun, currRun)

		// Update previous run to current, reset current run
		prevRun = currRun
		currRun = 1
	}

	// After the loop, add the last group
	result += min(prevRun, currRun)
	return result
}

func main() {
	// Example 1
	s1 := "00110011"
	expected1 := 6
	result1 := countBinarySubstrings(s1)
	println("Input:", s1)
	println("Expected Output:", expected1)
	println("Actual Output:", result1)
	println()

	// Example 2
	s2 := "10101"
	expected2 := 4
	result2 := countBinarySubstrings(s2)
	println("Input:", s2)
	println("Expected Output:", expected2)
	println("Actual Output:", result2)

	// Additional Test Case
	s3 := "00110"
	expected3 := 3
	result3 := countBinarySubstrings(s3)
	println("Input:", s3)
	println("Expected Output:", expected3)
	println("Actual Output:", result3)
}
