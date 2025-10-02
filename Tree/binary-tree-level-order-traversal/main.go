package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Define a level order list
	levelOrderList := [][]int{}
	// Initialize the queue with the root node
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// Number of nodes at the current level
		levelSize := len(queue)
		// Slice to store values of nodes at the current level
		levelOrder := []int{}

		// At each level, we'll iterate all elements at the level
		for i := 0; i < levelSize; i++ {
			curr := queue[0]  // Get the first node in the queue
			queue = queue[1:] // Pop the first element
			levelOrder = append(levelOrder, curr.Val)

			// Add left/right child to the queue if it exists
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}

		}

		// Add the current level's values to the result list
		levelOrderList = append(levelOrderList, levelOrder)
	}

	return levelOrderList
}

func main() {
	// Example 1: root = [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(root1))

	// Example 2: root = [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(levelOrder(root2))

	// Example 3: root = []
	var root3 *TreeNode = nil
	fmt.Println(levelOrder(root3))

	// Example 4: root = [1,2,3,4,null,null,5]
	root4 := &TreeNode{Val: 1}
	root4.Left = &TreeNode{Val: 2}
	root4.Right = &TreeNode{Val: 3}
	root4.Left.Left = &TreeNode{Val: 4}
	root4.Right.Right = &TreeNode{Val: 5}
	fmt.Println(levelOrder(root4)) // Expected: [[1],[2,3],[4,5]]
}
