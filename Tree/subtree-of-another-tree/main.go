package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree builds a binary tree from a slice (level order, using nil for missing nodes)
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i, node := range nodes {
		if node == nil {
			continue
		}
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < len(nodes) {
			node.Left = nodes[leftIdx]
		}
		if rightIdx < len(nodes) {
			node.Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func main() {
	// Example 1
	root1 := buildTree([]interface{}{3, 4, 5, 1, 2})
	subRoot1 := buildTree([]interface{}{4, 1, 2})
	fmt.Println("Example 1:")
	fmt.Println("Input: root = [3,4,5,1,2], subRoot = [4,1,2]")
	fmt.Println("Output:", isSubtree(root1, subRoot1))

	// Example 2
	root2 := buildTree([]interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0})
	subRoot2 := buildTree([]interface{}{4, 1, 2})
	fmt.Println("\nExample 2:")
	fmt.Println("Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]")
	fmt.Println("Output:", isSubtree(root2, subRoot2))
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Find a value in the root tree is similar to the root's value of the subRoot
	if root != nil && isSameTree(root, subRoot) {
		return true
	}
	// Find a value in the left' root tree is similar to the root's value of the subRoot
	if root.Left != nil && isSubtree(root.Left, subRoot) {
		return true
	}
	// Find a value in the right' root tree is similar to the root's value of the subRoot
	if root.Right != nil && isSubtree(root.Right, subRoot) {
		return true
	}

	return false
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
