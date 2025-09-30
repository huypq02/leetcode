package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// Define prev pointer with nil value
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // the next pointer
		curr.Next = prev  // point to the address of the previous pointer in the list
		prev = curr       // shift the previous pointer to the current pointer (start at head)
		curr = next
	}
	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	var list1, list2, prev *ListNode

	// Find middle of the linked list
	for fast != nil && fast.Next != nil {
		prev = slow
		// Shift the slow and fast pointers
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil // must set nil to avoid the infinite loop
	list1 = head
	list2 = reverseList(slow) // Reserve the list 2

	// Merge two sub-list (list1 and list2)
	tail := &ListNode{} // dummy for the tail of the list
	for list1 != nil && list2 != nil {

		// Merge node from list1
		tail.Next = list1
		list1 = list1.Next
		tail = tail.Next

		// Merge node from list2
		tail.Next = list2
		list2 = list2.Next
		tail = tail.Next
	}
}

// Helper function to create linked list from array
func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

// Helper function to print linked list
func printLinkedList(head *ListNode) {
	fmt.Print("[")
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(",")
		}
		current = current.Next
	}
	fmt.Println("]")
}

func main() {
	// Test Case 1: head = [1,2,3,4]
	// Expected Output: [1,4,2,3]
	fmt.Println("Test Case 1:")
	head1 := createLinkedList([]int{1, 2, 3, 4})
	fmt.Print("Input:  ")
	printLinkedList(head1)

	head1 = createLinkedList([]int{1, 2, 3, 4}) // Recreate since reorderList modifies the list
	reorderList(head1)
	fmt.Print("Output: ")
	printLinkedList(head1)
	fmt.Print("Expected: [1,4,2,3]")
	fmt.Println()
	fmt.Println()

	// Test Case 2: head = [1,2,3,4,5]
	// Expected Output: [1,5,2,4,3]
	fmt.Println("Test Case 2:")
	head2 := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("Input:  ")
	printLinkedList(head2)

	head2 = createLinkedList([]int{1, 2, 3, 4, 5}) // Recreate since reorderList modifies the list
	reorderList(head2)
	fmt.Print("Output: ")
	printLinkedList(head2)
	fmt.Print("Expected: [1,5,2,4,3]")
	fmt.Println()
}
