package main

import "fmt"

// Fixed-size sliding window
// Time complexity: O(n)
func decrypt(code []int, k int) []int {
	newCode := make([]int, len(code))
	maxLen := len(code)
	for right := 0; right < maxLen; right++ {
		if k > 0 {
			newCode[right] = sum(code, right+1, right+k)
		} else if k < 0 {
			newCode[right] = sum(code, maxLen+right+k, maxLen+right-1)
		} else {
			newCode[right] = 0
		}
	}
	return newCode
}

func sum(code []int, left, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += code[i%len(code)]
	}
	return sum
}

func main() {
	code := []int{5, 7, 1, 4}
	k := 3
	fmt.Println(decrypt(code, k))
	// Output: [12, 10, 16, 13]

	code = []int{2, 4, 9, 3}
	k = -2
	fmt.Println(decrypt(code, k))
	// Output: [12, 5, 6, 13]

	code = []int{1, 2, 3, 4}
	k = 0
	fmt.Println(decrypt(code, k))
	// Output: [0, 0, 0, 0]
}
