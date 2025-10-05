package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// Define a previous node with a minimum element
	// and a counter
	prev, kthSmallest, counter := math.MinInt64, math.MinInt64, 0

	// Define an inorder method
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		prev = node.Val // find the previous node with a minimum element
		counter++       // count to kth smallest element in a BST
		// When count is equal to k, the prev is kth smallest element in a BST
		if counter == k {
			kthSmallest = prev
			return
		}
		inorder(node.Right)
	}

	// Implement the inorder
	inorder(root)

	return kthSmallest
}

func main() {
	// Example 1: root = [3,1,4,null,2], k = 1
	// Tree structure:
	//     3
	//    / \
	//   1   4
	//    \
	//     2
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 2}
	k1 := 1
	fmt.Printf("Example 1: Output = %d (Expected: 1)\n", kthSmallest(root1, k1))

	// Example 2: root = [5,3,6,2,4,null,null,1], k = 3
	// Tree structure:
	//       5
	//      / \
	//     3   6
	//    / \
	//   2   4
	//  /
	// 1
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 6}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 1}
	k2 := 3
	fmt.Printf("Example 2: Output = %d (Expected: 3)\n", kthSmallest(root2, k2))
}
