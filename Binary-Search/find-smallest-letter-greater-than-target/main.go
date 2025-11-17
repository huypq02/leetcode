package main

func nextGreatestLetter(letters []byte, target byte) byte {
	// Initialize binary search boundaries
	left, right := 0, len(letters)-1
	// Default result is the first letter (for wrap-around case)
	result := letters[0]

	// Perform binary search to find the smallest letter greater than target
	for left <= right {
		// Calculate the middle index
		mid := left + (right-left)/2
		// If the middle letter is greater than target, it is a candidate
		if letters[mid] > target {
			// Update result and search left half for a smaller valid letter
			result = letters[mid]
			right = mid - 1
		} else {
			// Otherwise, search right half
			left = mid + 1
		}
	}

	// Return the result (either the found letter or the first letter if none is greater)
	return result
}

func main() {
	// Example 1
	letters1 := []byte{'c', 'f', 'j'}
	target1 := byte('a')
	// Expected output: 'c'
	println("Test 1 Output:", string(nextGreatestLetter(letters1, target1)), "| Expected: c")

	// Example 2
	letters2 := []byte{'c', 'f', 'j'}
	target2 := byte('c')
	// Expected output: 'f'
	println("Test 2 Output:", string(nextGreatestLetter(letters2, target2)), "| Expected: f")

	// Example 3
	letters3 := []byte{'x', 'x', 'y', 'y'}
	target3 := byte('z')
	// Expected output: 'x'
	println("Test 3 Output:", string(nextGreatestLetter(letters3, target3)), "| Expected: x")
}
