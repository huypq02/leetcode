package main

import "fmt"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Append a new value to the slice (dynamic array)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1] // Truncate the slice by reducing its length by one element
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1] // Top is the last element of the slice
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	// Map from value in nums2 to its next greater element.
	nextGreaterElemMap := make(map[int]int)

	// Stack to hold indices of nums2 elements whose next greater element hasn't been found yet.
	// Monotonic decreasing: ensures that we always compare current element to the top of stack.
	nextGreaterElemIdx := Stack{}
	for i := 0; i < len(nums2); i++ {
		// Pop elements from the stack while the current nums2[i] is greater than top of stack.
		// This means we found the next greater element for all popped elements.
		for len(nextGreaterElemIdx) > 0 && nums2[nextGreaterElemIdx.Top()] <= nums2[i] {
			prev := nextGreaterElemIdx.Top()
			nextGreaterElemIdx.Pop()
			nextGreaterElemMap[nums2[prev]] = nums2[i] // Record next greater for nums2[prev]
		}

		// Push current index onto stack for future comparisons.
		nextGreaterElemIdx.Push(i)
	}

	// Prepare result array for nums1 queries.
	// For each value in nums1, fetch its next greater from the map if exists, else -1.
	nextGreaterElem := make([]int, len(nums1))
	for i, v := range nums1 {
		if next, ok := nextGreaterElemMap[v]; ok {
			nextGreaterElem[i] = next
			continue
		}
		// If no next greater found, assign -1 as per problem statement.
		nextGreaterElem[i] = -1
	}

	return nextGreaterElem
}

func main() {
	// Example 1
	nums1a := []int{4, 1, 2}
	nums2a := []int{1, 3, 4, 2}
	// Expected output: [-1, 3, -1]
	result1 := nextGreaterElement(nums1a, nums2a)
	fmt.Println("Example 1 result:", result1, "Expected: [-1, 3, -1]")

	// Example 2
	nums1b := []int{2, 4}
	nums2b := []int{1, 2, 3, 4}
	// Expected output: [3, -1]
	result2 := nextGreaterElement(nums1b, nums2b)
	fmt.Println("Example 2 result:", result2, "Expected: [3, -1]")
}
