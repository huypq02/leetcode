package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to handle edge cases like removing the head
	dummy := &ListNode{Next: head}

	// Declare slow and fast pointers
	slow, fast := dummy, dummy

	// Define a counter
	i := 1
	// Find the nth node from the end of the list
	for fast.Next != nil {
		// Ensure fast stays n nodes ahead of slow
		if i <= n {
			fast = fast.Next
			i++
			continue
		}
		// Shift the slow and fast pointers when having a fixed window size
		slow = slow.Next
		fast = fast.Next
	}

	// the slow is the nth node from the end of the list
	// Skip the nth node from the end of the list
	slow.Next = slow.Next.Next

	return dummy.Next
}

// createList builds a linked list from a slice of integers
func createList(values []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, val := range values {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}

	return dummy.Next
}

// listToSlice converts a linked list to a slice for easy printing
func listToSlice(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func main() {
	// Example 1: head = [1,2,3,4,5], n = 2, expected output: [1,2,3,5]
	head1 := createList([]int{1, 2, 3, 4, 5})
	result1 := removeNthFromEnd(head1, 2)
	fmt.Printf("Example 1: Input = [1,2,3,4,5], n = 2, expected output: [1,2,3,5]\n")
	fmt.Printf("Output: %v\n\n", listToSlice(result1))

	// Example 2: head = [1], n = 1, expected output: []
	head2 := createList([]int{1})
	result2 := removeNthFromEnd(head2, 1)
	fmt.Printf("Example 2: Input = [1], n = 1, expected output: []\n")
	fmt.Printf("Output: %v\n\n", listToSlice(result2))

	// Example 3: head = [1,2], n = 1, expected output: [1]
	head3 := createList([]int{1, 2})
	result3 := removeNthFromEnd(head3, 1)
	fmt.Printf("Example 3: Input = [1,2], n = 1, expected output: [1]\n")
	fmt.Printf("Output: %v\n\n", listToSlice(result3))

	// Example 4: head = [1,2], n = 2, expected output: [2]
	head4 := createList([]int{1, 2})
	result4 := removeNthFromEnd(head4, 2)
	fmt.Printf("Example 4: Input = [1,2], n = 2, expected output: [2]\n")
	fmt.Printf("Output: %v\n\n", listToSlice(result4))
}
