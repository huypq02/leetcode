package main

import "fmt"

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum, currVal, carry := &ListNode{}, 0, 0
	dummy := sum
	for l1 != nil || l2 != nil {

		if l1 == nil {
			currVal = l2.Val + carry
		} else if l2 == nil {
			currVal = l1.Val + carry
		} else {
			currVal = l1.Val + l2.Val + carry
		}

		dummy.Next = &ListNode{Val: currVal % 10}
		carry = currVal / 10

		// Shift the left and right pointers
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		dummy = dummy.Next
	}

	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
		dummy = dummy.Next
	}

	return sum.Next
}

// Helper function to create a linked list from slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// Helper function to print linked list
func printLinkedList(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(",")
		}
		head = head.Next
	}
	fmt.Print("]")
}

func main() {
	// Test Case 1: l1 = [2,4,3], l2 = [5,6,4]
	// Expected output: [7,0,8] (342 + 465 = 807)
	fmt.Println("Test Case 1:")
	l1_1 := createLinkedList([]int{2, 4, 3})
	l2_1 := createLinkedList([]int{5, 6, 4})
	fmt.Print("Input: l1 = ")
	printLinkedList(l1_1)
	fmt.Print(", l2 = ")
	printLinkedList(l2_1)
	fmt.Println()
	fmt.Print("Output: ")
	result1 := addTwoNumbers(l1_1, l2_1)
	printLinkedList(result1)
	fmt.Println(" (Expected: [7,0,8])")
	fmt.Println()

	// Test Case 2: l1 = [0], l2 = [0]
	// Expected output: [0]
	fmt.Println("Test Case 2:")
	l1_2 := createLinkedList([]int{0})
	l2_2 := createLinkedList([]int{0})
	fmt.Print("Input: l1 = ")
	printLinkedList(l1_2)
	fmt.Print(", l2 = ")
	printLinkedList(l2_2)
	fmt.Println()
	fmt.Print("Output: ")
	result2 := addTwoNumbers(l1_2, l2_2)
	printLinkedList(result2)
	fmt.Println(" (Expected: [0])")
	fmt.Println()

	// Test Case 3: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	// Expected output: [8,9,9,9,0,0,0,1]
	fmt.Println("Test Case 3:")
	l1_3 := createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2_3 := createLinkedList([]int{9, 9, 9, 9})
	fmt.Print("Input: l1 = ")
	printLinkedList(l1_3)
	fmt.Print(", l2 = ")
	printLinkedList(l2_3)
	fmt.Println()
	fmt.Print("Output: ")
	result3 := addTwoNumbers(l1_3, l2_3)
	printLinkedList(result3)
	fmt.Println(" (Expected: [8,9,9,9,0,0,0,1])")
	fmt.Println()

	// Test Case 4: l1 = [3,7], l2 = [9,2]
	// Expected output: [2,0,1] (73 + 29 = 102)
	fmt.Println("Test Case 4:")
	l1_4 := createLinkedList([]int{3, 7})
	l2_4 := createLinkedList([]int{9, 2})
	fmt.Print("Input: l1 = ")
	printLinkedList(l1_4)
	fmt.Print(", l2 = ")
	printLinkedList(l2_4)
	fmt.Println()
	fmt.Print("Output: ")
	result4 := addTwoNumbers(l1_4, l2_4)
	printLinkedList(result4)
	fmt.Println(" (Expected: [2,0,1])")
}
