package main

import "fmt"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Push index onto stack
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1] // Pop index from stack
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1] // Get top index from stack
}

func largestRectangleArea(heights []int) int {
	stack := Stack{}
	maxArea := 0
	// Append a sentinel zero-height bar to ensure all bars in the stack are processed
	heights = append(heights, 0)
	for i, h := range heights {
		// While the current bar is shorter than the bar at the top of the stack,
		// calculate area for the bar at the top of the stack
		for len(stack) > 0 && h < heights[stack.Top()] {
			prevIdx := stack.Top()
			stack.Pop()
			height := heights[prevIdx]
			// If the stack is empty, the rectangle extends from index 0 to i-1
			// Otherwise, it's blocked on the left by stack.Top()
			width := i
			if len(stack) > 0 {
				width = i - stack.Top() - 1 // Notice the current stack.Top()
			}
			if maxArea < width*height {
				maxArea = width * height
			}
		}
		// Push the current index if the height is greater than or equal to the bar at the top of the stack
		// This maintains a stack of indices with non-decreasing heights
		stack.Push(i)
	}
	return maxArea
}

func main() {
	// Test Case 1: Example 1
	heights1 := []int{2, 1, 5, 6, 2, 3}
	result1 := largestRectangleArea(heights1)
	fmt.Printf("Test Case 1:\n")
	fmt.Printf("Input: heights = %v\n", heights1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 10\n\n")

	// Test Case 2: Example 2
	heights2 := []int{2, 4}
	result2 := largestRectangleArea(heights2)
	fmt.Printf("Test Case 2:\n")
	fmt.Printf("Input: heights = %v\n", heights2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 4\n\n")
}
