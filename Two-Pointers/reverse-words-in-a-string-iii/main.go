package main

func reverseWords(s string) string {
	return ""
}

func main() {
	// Example 1
	input1 := "Let's take LeetCode contest"
	expected1 := "s'teL ekat edoCteeL tsetnoc"
	result1 := reverseWords(input1)
	println("Input:", input1)
	println("Expected Output:", expected1)
	println("Actual Output:", result1)
	println()

	// Example 2
	input2 := "Mr Ding"
	expected2 := "rM gniD"
	result2 := reverseWords(input2)
	println("Input:", input2)
	println("Expected Output:", expected2)
	println("Actual Output:", result2)
}
