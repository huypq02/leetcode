package main

func duplicateZeros(arr []int) {
	n := len(arr)
	readPos := 0
	dest := 0

	// First pass: Find the position up to which we need to process (simulate the final array size after duplicating zeros)
	// dest: the position in the final array (could be > n if zeros are duplicated)
	// readPos: the position in the original array
	for dest < n {
		if arr[readPos] == 0 {
			// If current element is zero, it will be duplicated, so move dest one extra step
			dest++
		}
		dest++
		readPos++
	}

	// If readPos == n, it means the last element fits exactly, no need to shift
	if readPos == n {
		return
	}

	// Second pass: Work backwards to fill the array in place
	writePos := n - 1
	readPos-- // Set readPos to the last element to process

	// If dest > n, the last zero needs to be handled specially (it was partially duplicated)
	if dest > n && arr[readPos] == 0 {
		arr[writePos] = arr[readPos]
		writePos--
		readPos--
	}

	// Fill the array from the end, duplicating zeros as needed
	for writePos > 0 {
		if writePos > 0 && arr[readPos] == 0 {
			// Duplicate zero
			arr[writePos] = arr[readPos]
			writePos--
		}
		// Copy current element
		arr[writePos] = arr[readPos]
		writePos--
		readPos--
	}
}

func main() {
	// Example 1
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	// Expected output after duplicateZeros: [1,0,0,2,3,0,0,4]
	println("Test case 1 input: 1 0 2 3 0 4 5 0")
	duplicateZeros(arr1)
	println("Test case 1 result (expected: 1 0 0 2 3 0 0 4):")
	for _, v := range arr1 {
		print(v, " ")
	}
	println()

	// Example 2
	arr2 := []int{1, 2, 3}
	// Expected output after duplicateZeros: [1,2,3]
	println("Test case 2 input: 1 2 3")
	duplicateZeros(arr2)
	println("Test case 2 result (expected: 1 2 3):")
	for _, v := range arr2 {
		print(v, " ")
	}
	println()

	// Example 3
	arr3 := []int{8, 4, 5, 0, 0, 0, 0, 7}
	// Expected output after duplicateZeros: [8,4,5,0,0,0,0,0]
	println("Test case 3 input: 8 4 5 0 0 0 0 7")
	duplicateZeros(arr3)
	println("Test case 3 result (expected: 8 4 5 0 0 0 0 0):")
	for _, v := range arr3 {
		print(v, " ")
	}
	println()
}
