package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	isValid := true // Set flag is true
	var dfs func(node *TreeNode, min, max int)
	dfs = func(node *TreeNode, min, max int) {
		if node == nil {
			return
		}

		// In this case, it is not a valid binary search tree
		if min >= node.Val || max <= node.Val {
			isValid = false
			return
		}

		// Check the left subtrees of the current node
		dfs(node.Left, min, node.Val)
		// Check the right subtrees of the current node
		dfs(node.Right, node.Val, max)

	}

	// Implement dfs with MinInt64 < root.Val < MaxInt64
	dfs(root, math.MinInt64, math.MaxInt64)

	return isValid
}

func main() {
	// Example 1: root = [2,1,3]
	// Tree structure:
	//     2
	//    / \
	//   1   3
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Example 1:")
	fmt.Printf("Input: root = [2,1,3]\n")
	fmt.Printf("Output: %v; Expected Output: true\n", isValidBST(root1))

	// Example 2: root = [5,1,4,null,null,3,6]
	// Tree structure:
	//     5
	//    / \
	//   1   4
	//      / \
	//     3   6
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}
	fmt.Println("\nExample 2:")
	fmt.Printf("Input: root = [5,1,4,null,null,3,6]\n")
	fmt.Printf("Output: %v; Expected Output: false\nExplanation: The root node's value is 5 but its right child's value is 4.\n", isValidBST(root2))

	// Example 3: root = [5,4,6,null,null,3,7]
	// Tree structure:
	//     5
	//    / \
	//   4   6
	//      / \
	//     3   7
	root3 := &TreeNode{Val: 5}
	root3.Left = &TreeNode{Val: 4}
	root3.Right = &TreeNode{Val: 6}
	root3.Right.Left = &TreeNode{Val: 3}
	root3.Right.Right = &TreeNode{Val: 7}
	fmt.Println("\nExample 3:")
	fmt.Printf("Input: root = [5,4,6,null,null,3,7]\n")
	fmt.Printf("Output: %v; Expected Output: false\n", isValidBST(root3))

	// Example 4: root = [0]
	// Tree structure:
	//     0
	root4 := &TreeNode{Val: 0}
	fmt.Println("\nExample 4:")
	fmt.Printf("Input: root = [0]\n")
	fmt.Printf("Output: %v; Expected Output: true\n", isValidBST(root4))

	// Example 5: root = [3,1,5,0,2,4,6]
	// Tree structure:
	//         3
	//       /   \
	//      1     5
	//     / \   / \
	//    0   2 4   6
	root5 := &TreeNode{Val: 3}
	root5.Left = &TreeNode{Val: 1}
	root5.Right = &TreeNode{Val: 5}
	root5.Left.Left = &TreeNode{Val: 0}
	root5.Left.Right = &TreeNode{Val: 2}
	root5.Right.Left = &TreeNode{Val: 4}
	root5.Right.Right = &TreeNode{Val: 6}
	fmt.Println("\nExample 5:")
	fmt.Printf("Input: root = [3,1,5,0,2,4,6]\n")
	fmt.Printf("Output: %v; Expected Output: true\n", isValidBST(root5))
}
