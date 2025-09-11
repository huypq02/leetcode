package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createListCycle(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	prevHead := &ListNode{}
	curr := prevHead
	var nextTail *ListNode
	for i, v := range nums {
		// Define the address of the next tail
		if i == pos+1 {
			nextTail = curr
		}

		curr.Next = &ListNode{Val: v}

		// Assign the address of the next tail
		if i == len(nums)-1 {
			curr = nextTail
			break // End
		}
		curr = curr.Next
	}
	return prevHead.Next // the head of the linked-list
}

func printList(list *ListNode) {
	curr := list
	for curr != nil {
		fmt.Print(curr.Val, " ")
		fmt.Print(curr, " ")

		curr = curr.Next
	}
	fmt.Println()
}

func hasCycle(head *ListNode) bool {
	address := make(map[*ListNode]bool)
	current := head

	for current != nil {
		if ok := address[current]; ok {
			return true
		}
		address[current] = true
	}
	return false
}

func main() {

	// Input: head = [3,2,0,-4], pos = 1
	head1 := createListCycle([]int{3, 2, 0, -4}, 1)
	// Output: true
	println(hasCycle(head1))

	// Input: head = [1,2], pos = 0
	// Output: true

	// Input: head = [1], pos = -1
	// Output: false

}
