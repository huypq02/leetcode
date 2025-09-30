package main

import "fmt"

// Node represents a node in the linked list with random pointer
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var newNode *Node
	oldToNew := make(map[*Node]*Node)
	current := head

	for current != nil {
		// New map of current node (old node)
		oldToNew[current] = &Node{Val: current.Val}
		newNode = oldToNew[current]

		// Move to next position
		newNode = newNode.Next
		current = current.Next
	}

	current = head
	newNode = oldToNew[current]
	for current != nil {
		// Assign the value of the current node to the new node
		newNode.Next = oldToNew[current.Next]
		newNode.Random = oldToNew[current.Random]

		// Move to next position
		newNode = newNode.Next
		current = current.Next
	}

	newNode = oldToNew[head]
	return newNode
}

// Helper function to create a linked list from array representation
func createLinkedList(values [][]interface{}) *Node {
	if len(values) == 0 {
		return nil
	}

	// Create all nodes first
	nodes := make([]*Node, len(values))
	for i, val := range values {
		nodes[i] = &Node{Val: val[0].(int)}
	}

	// Link next pointers
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Link random pointers
	for i, val := range values {
		if val[1] != nil {
			randomIndex := val[1].(int)
			nodes[i].Random = nodes[randomIndex]
		}
	}

	return nodes[0]
}

// Helper function to convert linked list to array representation for testing
func linkedListToArray(head *Node) [][]interface{} {
	if head == nil {
		return [][]interface{}{}
	}

	// First pass: collect all nodes and create index map
	nodeToIndex := make(map[*Node]int)
	current := head
	index := 0
	for current != nil {
		nodeToIndex[current] = index
		current = current.Next
		index++
	}

	// Second pass: build result array
	result := make([][]interface{}, 0)
	current = head
	for current != nil {
		var randomIndex interface{}
		if current.Random != nil {
			randomIndex = nodeToIndex[current.Random]
		}
		result = append(result, []interface{}{current.Val, randomIndex})
		current = current.Next
	}

	return result
}

func main() {
	// Test Case 1: [[7,null],[13,0],[11,4],[10,2],[1,0]]
	fmt.Println("Test Case 1:")
	input1 := [][]interface{}{
		{7, nil},
		{13, 0},
		{11, 4},
		{10, 2},
		{1, 0},
	}
	head1 := createLinkedList(input1)
	copied1 := copyRandomList(head1)
	result1 := linkedListToArray(copied1)
	fmt.Printf("Input:  %v\n", input1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Println()

	// Test Case 2: [[1,1],[2,1]]
	fmt.Println("Test Case 2:")
	input2 := [][]interface{}{
		{1, 1},
		{2, 1},
	}
	head2 := createLinkedList(input2)
	copied2 := copyRandomList(head2)
	result2 := linkedListToArray(copied2)
	fmt.Printf("Input:  %v\n", input2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Println()

	// Test Case 3: [[3,null],[3,0],[3,null]]
	fmt.Println("Test Case 3:")
	input3 := [][]interface{}{
		{3, nil},
		{3, 0},
		{3, nil},
	}
	head3 := createLinkedList(input3)
	copied3 := copyRandomList(head3)
	result3 := linkedListToArray(copied3)
	fmt.Printf("Input:  %v\n", input3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Println()

	// Test Case 4: Empty list
	fmt.Println("Test Case 4 (Empty list):")
	var head4 *Node = nil
	copied4 := copyRandomList(head4)
	result4 := linkedListToArray(copied4)
	fmt.Printf("Input:  []\n")
	fmt.Printf("Output: %v\n", result4)
}
