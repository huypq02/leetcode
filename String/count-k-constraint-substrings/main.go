package main

import "fmt"

// Not using window sliding technique
// Time complexity: O(n^2)
func countKConstraintSubstrings00(s string, k int) int {
	count := 0
	n := len(s)

	for i := 0; i < n; i++ {
		zeroCount := 0
		oneCount := 0

		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeroCount++
			} else {
				oneCount++
			}

			// Check if constraint is satisfied
			if zeroCount <= k || oneCount <= k {
				count++
			} else {
				// Break early if constraint is violated
				break
			}
		}
	}
	return count
}

// Using window sliding technique
// Time complexity: O(n)
func countKConstraintSubstrings01(s string, k int) int {
	n := len(s)
	left := 0
	currentOnes := 0
	currentZeros := 0
	count := 0

	for right := 0; right < n; right++ {
		if s[right] == '1' {
			currentOnes++
		} else {
			currentZeros++
		}

		for currentOnes > k && currentZeros > k {
			if s[left] == '1' {
				currentOnes--
			} else {
				currentZeros--
			}
			left++
		}
		count += right - left + 1
	}

	return count
}

func main() {
	// Test case 1
	s1 := "10101"
	k1 := 1
	fmt.Println("Test case 1 does not use sliding technique: ", countKConstraintSubstrings00(s1, k1))
	fmt.Println("Test case 1 uses sliding window technique: ", countKConstraintSubstrings01(s1, k1))

	// Test case 2
	s2 := "1010101"
	k2 := 2
	fmt.Println("Test case 2 does not use sliding technique: ", countKConstraintSubstrings00(s2, k2))
	fmt.Println("Test case 2 uses sliding window technique: ", countKConstraintSubstrings01(s2, k2))
}
