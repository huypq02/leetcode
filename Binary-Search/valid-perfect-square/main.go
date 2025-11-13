package main

func isPerfectSquare(num int) bool {
	left, right := 0, num
	// Perform binary search between left and right
	for left <= right {
		mid := left + (right-left)/2 // Avoids overflow
		// If mid*mid is less than num, search the right half
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			// If mid*mid is greater than num, search the left half
			right = mid - 1
		} else {
			// Found an integer whose square is num
			return true
		}
	}
	// No integer square root found
	return false
}

func main() {
	// Example 1:
	num1 := 16
	expected1 := true
	result1 := isPerfectSquare(num1)
	println("Input:", num1, "Expected:", expected1, "Output:", result1)

	// Example 2:
	num2 := 14
	expected2 := false
	result2 := isPerfectSquare(num2)
	println("Input:", num2, "Expected:", expected2, "Output:", result2)
}
