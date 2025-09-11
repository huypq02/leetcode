package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var max int = 0

	if root == nil {
		return 0
	}
	max++

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		max += left
	} else {
		max += right
	}

	return max
}

func main() {
	// Input: root = [3,9,20,null,null,15,7]
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{9, nil, nil},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{15, nil, nil},
			Right: &TreeNode{7, nil, nil},
		},
	}
	// Output: 3
	println(maxDepth(tree))

	// Input: root = [1,null,2]
	tree2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	// Output: 2
	println(maxDepth(tree2))
}
