package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// }

func main() {
	l1 := ListNode{2, nil}
	l2 := ListNode{4, &ListNode{5, &ListNode{6, nil}}}
	l1.Next = &ListNode{4, nil}
	fmt.Println(l1)
	current := l1.Val
	fmt.Println(current)
	for l2.Next != nil {
		fmt.Println(l2.Val)
		l2 = *l2.Next
	}
	fmt.Println(l2.Val)
}
