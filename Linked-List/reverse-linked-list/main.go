package main

import "fmt"

//  Definition for singly-linked list.
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
	head = prev
	return head
}

func createLinkedList(num []int) *ListNode {
	if len(num) == 0 {
		return nil
	}
	head := &ListNode{Val: num[0]}
	current := head
	for _, v := range num[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func printList(list *ListNode) {
	curr := list
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
	fmt.Println()
}

func main() {

	// Input: head = [1,2,3,4,5]
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	// Output: [5,4,3,2,1]
	printList(reverseList(head))

	// Input: head = [1,2]
	head2 := createLinkedList([]int{1, 2})
	// Output: [2,1]
	printList(reverseList(head2))

	// Input: head = []
	head3 := createLinkedList([]int{})
	// Output: []
	printList(reverseList(head3))

}
