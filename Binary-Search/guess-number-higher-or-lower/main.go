package main

import "fmt"

// guess API as required by the exercise
var pick int

func guess(num int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {

	// Binary search to find the picked number
	left, right := 1, n
	for left <= right {
		// Calculate the middle value to avoid integer overflow
		mid := left + (right-left)/2
		if guess(mid) == 1 {
			// The picked number is higher than mid, search right half
			left = mid + 1
		} else if guess(mid) == -1 {
			// The picked number is lower than mid, search left half
			right = mid - 1
		} else {
			// Found the picked number
			return mid
		}
	}

	return 0
}

func main() {
	// Test case 1: n = 10, pick = 6
	pick = 6
	n := 10
	fmt.Printf("Input: n = %d, pick = %d\n", n, pick)
	fmt.Printf("Output: %d\n", guessNumber(n))
	fmt.Println("Expected Output: 6")

	// Test case 2: n = 1, pick = 1
	pick = 1
	n = 1
	fmt.Printf("Input: n = %d, pick = %d\n", n, pick)
	fmt.Printf("Output: %d\n", guessNumber(n))
	fmt.Println("Expected Output: 1")

	// Test case 3: n = 2, pick = 1
	pick = 1
	n = 2
	fmt.Printf("Input: n = %d, pick = %d\n", n, pick)
	fmt.Printf("Output: %d\n", guessNumber(n))
	fmt.Println("Expected Output: 1")
}
