package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calculateDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth - 1
	}

	left := calculateDepth(node.Left, depth+1)
	right := calculateDepth(node.Right, depth+1)

	if left-right > 1 || right-left > 1 {
		return -1
	}
	println(left, ":", right)

	// Maximum of depth
	return max(left, right)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if calculateDepth(root, 1) == -1 {
		return false
	}
	return true

}

func main() {
	// Input: root = [3,9,20,null,null,15,7]
	// Output: true
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{9, nil, nil},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{15, nil, nil},
			Right: &TreeNode{7, nil, nil},
		},
	}
	println(isBalanced(tree))

	// Input: root = [1,2,2,3,3,null,null,4,4]
	// Output: false
	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{4, nil, nil},
				Right: &TreeNode{4, nil, nil},
			},
			Right: &TreeNode{3, nil, nil},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	println(isBalanced(tree2))

	// Input: root = []
	// Output: true
	tree3 := &TreeNode{}
	println(isBalanced(tree3))

	// Input: root = [1,null,2,null,3]
	// Output: false
	tree4 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{3, nil, nil},
		},
	}
	println(isBalanced(tree4))
}


