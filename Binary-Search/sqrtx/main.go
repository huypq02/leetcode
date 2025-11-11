package main

func mySqrt(x int) int {
	// Binary search approach to find the integer square root of x
	// The answer is the largest integer whose square is less than or equal to x
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2 // Avoids overflow
		// If mid*mid is less than x, move to the right half
		if x > mid*mid {
			left = mid + 1
		} else if x < mid*mid {
			// If mid*mid is greater than x, move to the left half
			right = mid - 1
		} else {
			// If mid*mid equals x, mid is the answer
			return mid
		}
	}
	// When the loop ends, right is the integer part of the square root
	return right
}

func main() {
	// Example 1:
	x1 := 4
	expected1 := 2
	result1 := mySqrt(x1)
	println("Input:", x1, "Expected Output:", expected1, "Actual Output:", result1)

	// Example 2:
	x2 := 8
	expected2 := 2
	result2 := mySqrt(x2)
	println("Input:", x2, "Expected Output:", expected2, "Actual Output:", result2)

	// Example 3:
	x3 := 1
	expected3 := 1
	result3 := mySqrt(x3)
	println("Input:", x3, "Expected Output:", expected3, "Actual Output:", result3)
}
