package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var rightSideView = []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var lastNodeAtLevel int

		for i := 0; i < levelSize; i++ {
			curr := queue[0]  // Get the first node
			queue = queue[1:] // Pop the first node of queue

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}

			// Last node at the end of the current level
			// will be the value of the node on the right side view
			if i == levelSize-1 {
				lastNodeAtLevel = curr.Val
			}
		}

		// Add the value of the node on the right side view
		rightSideView = append(rightSideView, lastNodeAtLevel)
	}
	return rightSideView
}

func main() {
	// Example 1: root = [1,2,3,null,5,null,4]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	fmt.Println("Example 1 Output:", rightSideView(root1), "Expected: [1,3,4]")

	// Example 2: root = [1,2,3,4,null,null,null,5]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("Example 2 Output:", rightSideView(root2), "Expected: [1,3,4,5]")

	// Example 3: root = [1,null,3]
	root3 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("Example 3 Output:", rightSideView(root3), "Expected: [1,3]")

	// Example 4: root = []
	var root4 *TreeNode = nil
	fmt.Println("Example 4 Output:", rightSideView(root4), "Expected: []")
}
