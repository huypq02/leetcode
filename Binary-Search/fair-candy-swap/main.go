package main

import "sort"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// The goal is to find a pair (a, b) such that:
	// sumA - a + b == sumB - b + a
	// Rearranged: sumA - sumB = 2a - 2b => (sumA - sumB) / 2 = a - b
	// So, for each a in aliceSizes, we want to find b in bobSizes such that:
	// a = b + (sumA - sumB) / 2

	// Calculate the sum of Alice's and Bob's candies
	sumA, sumB := 0, 0
	for _, v := range aliceSizes {
		sumA += v
	}
	for _, v := range bobSizes {
		sumB += v
	}

	// Compute the difference divided by 2
	diff := (sumA - sumB) / 2 // This is a - b

	// Sort Bob's sizes to enable binary search
	sort.Ints(bobSizes)

	// For each candy box size Alice has, calculate the value Bob should have
	for i := 0; i < len(aliceSizes); i++ {
		// We want to find b in bobSizes such that a = b + diff, so b = a - diff
		bExpected := aliceSizes[i] - diff
		// Binary search for bExpected in bobSizes
		left, right := 0, len(bobSizes)-1
		for left <= right {
			mid := left + (right-left)/2
			if bExpected > bobSizes[mid] {
				left = mid + 1
			} else if bExpected < bobSizes[mid] {
				right = mid - 1
			} else {
				// Found the pair (a, b) that satisfies the condition
				return []int{aliceSizes[i], bobSizes[mid]}
			}
		}
	}

	// If no valid pair is found (though the problem guarantees one exists)
	return []int{}
}

func main() {
	// Example 1
	alice := []int{1, 1}
	bob := []int{2, 2}
	result := fairCandySwap(alice, bob)
	println("Example 1 result:", result[0], result[1], "(expected: 1 2)")

	// Example 2
	alice = []int{1, 2}
	bob = []int{2, 3}
	result = fairCandySwap(alice, bob)
	println("Example 2 result:", result[0], result[1], "(expected: 1 2)")

	// Example 3
	alice = []int{2}
	bob = []int{1, 3}
	result = fairCandySwap(alice, bob)
	println("Example 3 result:", result[0], result[1], "(expected: 2 3)")
}
