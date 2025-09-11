package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// Validate before merging the two
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	p1, p2 := list1, list2
	// Declare a result
	var result *ListNode
	if p1.Val < p2.Val {
		result = &ListNode{Val: p1.Val}
		p1 = p1.Next
	} else {
		result = &ListNode{Val: p2.Val}
		p2 = p2.Next
	}

	currResult := result
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			currResult.Next = &ListNode{Val: p1.Val}
			p1 = p1.Next
			currResult = currResult.Next
		} else {
			currResult.Next = &ListNode{Val: p2.Val}
			p2 = p2.Next
			currResult = currResult.Next
		}
	}

	for p1 != nil {
		currResult.Next = &ListNode{Val: p1.Val}
		p1 = p1.Next
		currResult = currResult.Next
	}

	for p2 != nil {
		currResult.Next = &ListNode{Val: p2.Val}
		p2 = p2.Next
		currResult = currResult.Next
	}

	return result
}

func createList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	var head = &ListNode{Val: list[0]} // define the head of the linked list
	var curr = head
	// Starting the index of list is 1
	for _, v := range list[1:] {
		curr.Next = &ListNode{Val: v} // set val for the next pointer
		curr = curr.Next              // set the address of the next pointer
	}
	return head
}

func printList(head *ListNode) {
	var curr = head
	for curr != nil {
		print(curr.Val, " ")
		curr = curr.Next
	}
	println()
}

func main() {
	// Input: list1 = [1,2,4], list2 = [1,3,4]
	list1, list2 := createList([]int{1, 2, 4}), createList([]int{1, 3, 4})
	// Output: [1,1,2,3,4,4]
	printList(mergeTwoLists(list1, list2))

	// Input: list1 = [], list2 = []
	list1b, list2b := createList([]int{}), createList([]int{})
	// Output: []
	printList(mergeTwoLists(list1b, list2b))

	// Input: list1 = [], list2 = [0]
	list1c, list2c := createList([]int{}), createList([]int{0})
	// Output: [0]
	printList(mergeTwoLists(list1c, list2c))

}
