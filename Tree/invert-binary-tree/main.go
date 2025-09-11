package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}

func main() {
	// Input: root = [4,2,7,1,3,6,9]
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{1, nil, nil},
			Right: &TreeNode{3, nil, nil},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{6, nil, nil},
			Right: &TreeNode{9, nil, nil},
		},
	}
	// Output: [4,7,2,9,6,3,1]
	fmt.Println(invertTree(tree))

	// Input: root = [2,1,3]
	tree2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}
	// Output: [2,3,1]
	fmt.Println(invertTree(tree2))

	// Input: root = []
	tree3 := &TreeNode{}
	// Output: []
	fmt.Println(invertTree(tree3))
}
