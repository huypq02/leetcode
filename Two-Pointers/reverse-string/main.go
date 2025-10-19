package main

func reverseString(s []byte) {

	left, right := 0, len(s)-1
	for left < right {
		// Swap the elements at the left and right pointers
		temp := s[right]
		s[right] = s[left]
		s[left] = temp

		// Move the pointers towards each other
		left++  // Increment left pointer
		right-- // Decrement right pointer
	}
}

func main() {
	// Example 1
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
	println("Example 1 result:", string(s1), "| Expected: olleh")

	// Example 2
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s2)
	println("Example 2 result:", string(s2), "| Expected: hannaH")
}
