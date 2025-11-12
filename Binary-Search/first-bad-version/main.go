package main

func firstBadVersion(n int) int {
	// Initialize search boundaries
	left, right := 1, n
	// Variable to store the first bad version found
	firstBadVersion := n
	// Perform binary search
	for left <= right {
		// Find the middle version
		mid := left + (right-left)/2
		// If mid is not bad, search the right half
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			// If mid is bad, record it and search the left half
			firstBadVersion = mid
			right = mid - 1
		}
	}
	// Return the first bad version found
	return firstBadVersion
}

// Mocked isBadVersion for testing
var badVersion int

func isBadVersion(version int) bool {
	return version >= badVersion
}

func main() {
	// Example 1: n = 5, bad = 4
	badVersion = 4
	n := 5
	// Expected output: 4
	println("Test Case 1:")
	println("Input: n =", n, ", bad =", badVersion)
	println("Expected Output: 4")
	println("Actual Output:", firstBadVersion(n))
	println()

	// Example 2: n = 1, bad = 1
	badVersion = 1
	n = 1
	// Expected output: 1
	println("Test Case 2:")
	println("Input: n =", n, ", bad =", badVersion)
	println("Expected Output: 1")
	println("Actual Output:", firstBadVersion(n))
}
