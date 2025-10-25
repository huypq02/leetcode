package main

func reverseStr(s string, k int) string {
	return ""
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
}
