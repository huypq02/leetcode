package main

// isLetter checks if a character is an English letter (a-z or A-Z)
func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func reverseOnlyLetters(s string) string {
	// Convert string to byte slice since strings are immutable in Go
	b := []byte(s)

	// Initialize two pointers at both ends
	left, right := 0, len(b)-1

	// Continue until pointers meet
	for left < right {
		// Skip non-letter characters from the left
		if !isLetter(b[left]) {
			left++
			continue
		}
		// Skip non-letter characters from the right
		if !isLetter(b[right]) {
			right--
			continue
		}

		// Both are letters, swap them
		b[left], b[right] = b[right], b[left]

		// Move both pointers inward
		left++
		right--
	}

	// Convert byte slice back to string
	return string(b)
}

func main() {
	// Example 1
	input1 := "ab-cd"
	expected1 := "dc-ba"
	println("Input:", input1)
	println("Expected Output:", expected1)
	println("Actual Output:", reverseOnlyLetters(input1))
	println()

	// Example 2
	input2 := "a-bC-dEf-ghIj"
	expected2 := "j-Ih-gfE-dCba"
	println("Input:", input2)
	println("Expected Output:", expected2)
	println("Actual Output:", reverseOnlyLetters(input2))
	println()

	// Example 3
	input3 := "Test1ng-Leet=code-Q!"
	expected3 := "Qedo1ct-eeLg=ntse-T!"
	println("Input:", input3)
	println("Expected Output:", expected3)
	println("Actual Output:", reverseOnlyLetters(input3))
	println()

	// Example 4
	input4 := "7_28]"
	expected4 := "7_28]"
	println("Input:", input4)
	println("Expected Output:", expected4)
	println("Actual Output:", reverseOnlyLetters(input4))
	println()
}
