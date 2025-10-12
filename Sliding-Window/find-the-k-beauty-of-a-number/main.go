package main

import (
	"strconv"
)

// isDivisor checks if 'sub' is a divisor of 'num'.
func isDivisor(num, sub int) bool {
	return num%sub == 0
}

func divisorSubstrings(num int, k int) int {
	// Convert num to string to extract substrings of length k.
	numStr := strconv.Itoa(num)
	left, count := 0, 0

	// Slide a window of size k over numStr.
	for right := k - 1; right < len(numStr); right++ {
		subStr := numStr[left : right+1]

		// Convert substring to integer.
		subNum, err := strconv.Atoi(subStr)
		if err != nil {
			// If conversion fails, skip this substring.
			return 0
		}

		// Ignore substrings that are '0' (0 is not a divisor).
		if subNum == 0 {
			left++
			continue
		}

		// If subNum is a divisor of num, increment count.
		if isDivisor(num, subNum) {
			count++
		}

		left++
	}

	return count
}

func main() {
	// // Example 1: num = 240, k = 2, expected output: 2
	// num1 := 240
	// k1 := 2
	// expected1 := 2
	// result1 := divisorSubstrings(num1, k1)
	// println("Testcase 1:", "Input:", num1, k1, "Expected:", expected1, "Got:", result1)

	// // Example 2: num = 430043, k = 2, expected output: 2
	// num2 := 430043
	// k2 := 2
	// expected2 := 2
	// result2 := divisorSubstrings(num2, k2)
	// println("Testcase 2:", "Input:", num2, k2, "Expected:", expected2, "Got:", result2)

	// Additional Testcase: num = 604052, k = 1, expected output: 2
	num3 := 604052
	k3 := 1
	expected3 := 2
	result3 := divisorSubstrings(num3, k3)
	println("Testcase 3:", "Input:", num3, k3, "Expected:", expected3, "Got:", result3)
}
