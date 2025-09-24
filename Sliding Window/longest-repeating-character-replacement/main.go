package main

import "fmt"

func characterReplacement(s string, k int) int {

	mapSubstringFreq := make(map[byte]int)
	left, maxFreq, maxFreqInWindow := 0, 0, 0
	for right := 0; right < len(s); right++ {

		// Store the substring existed to the mapSubstringFreq map
		mapSubstringFreq[s[right]]++
		// Update max frequency in current window
		if maxFreqInWindow < mapSubstringFreq[s[right]] {
			maxFreqInWindow = mapSubstringFreq[s[right]]
		}

		// Shift the left pointer to the next index
		// if the distance between the right pointer and the window size is greater than k
		if (right-left+1)-maxFreqInWindow > k {
			mapSubstringFreq[s[left]]--
			left++
		}

		// Update max length with current valid window size
		if maxFreq < right-left+1 {
			maxFreq = right - left + 1
		}
	}

	return maxFreq
}

func main() {
	// Test case 1: Example 1 from README
	s1 := "ABAB"
	k1 := 2
	expected1 := 4
	result1 := characterReplacement(s1, k1)
	fmt.Printf("Test 1: s = %q, k = %d\n", s1, k1)
	fmt.Printf("Expected: %d, Got: %d\n", expected1, result1)
	if result1 == expected1 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Test case 2: Example 2 from README
	s2 := "AABABBA"
	k2 := 1
	expected2 := 4
	result2 := characterReplacement(s2, k2)
	fmt.Printf("Test 2: s = %q, k = %d\n", s2, k2)
	fmt.Printf("Expected: %d, Got: %d\n", expected2, result2)
	if result2 == expected2 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Additional test cases for edge cases
	// Test case 3: All same characters
	s3 := "AAAA"
	k3 := 0
	expected3 := 4
	result3 := characterReplacement(s3, k3)
	fmt.Printf("Test 3: s = %q, k = %d\n", s3, k3)
	fmt.Printf("Expected: %d, Got: %d\n", expected3, result3)
	if result3 == expected3 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Test case 4: Single character
	s4 := "A"
	k4 := 1
	expected4 := 1
	result4 := characterReplacement(s4, k4)
	fmt.Printf("Test 4: s = %q, k = %d\n", s4, k4)
	fmt.Printf("Expected: %d, Got: %d\n", expected4, result4)
	if result4 == expected4 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Test case 5: AAAB with k = 0
	s5 := "AAAB"
	k5 := 0
	expected5 := 3
	result5 := characterReplacement(s5, k5)
	fmt.Printf("Test 5: s = %q, k = %d\n", s5, k5)
	fmt.Printf("Expected: %d, Got: %d\n", expected5, result5)
	if result5 == expected5 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Test case 6: ABAA with k = 0
	s6 := "ABAA"
	k6 := 0
	expected6 := 2
	result6 := characterReplacement(s6, k6)
	fmt.Printf("Test 6: s = %q, k = %d\n", s6, k6)
	fmt.Printf("Expected: %d, Got: %d\n", expected6, result6)
	if result6 == expected6 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}

	// Test case 7: Long string with k = 7
	s7 := "EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH"
	k7 := 7
	expected7 := 11
	result7 := characterReplacement(s7, k7)
	fmt.Printf("Test 7: s = %q, k = %d\n", s7, k7)
	fmt.Printf("Expected: %d, Got: %d\n", expected7, result7)
	if result7 == expected7 {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}
}
