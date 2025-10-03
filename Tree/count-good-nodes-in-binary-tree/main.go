package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	count := 0
	var dfs func(node *TreeNode, max int)
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}

		// Check the current value is greater than
		// or equal to the maximum value of the parent nodes.
		if node.Val >= max {
			max = node.Val
			count++
		}

		// Check the left of the current node
		dfs(node.Left, max)
		// Check the right of the current node
		dfs(node.Right, max)
	}

	// Implement dfs
	dfs(root, root.Val)

	return count
}

func main() {
	// Example 1: root = [3,1,4,3,null,1,5]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}
	root1.Left.Left = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 1}
	root1.Right.Right = &TreeNode{Val: 5}
	// Expected output: 4
	println("Example 1 Output:", goodNodes(root1), "| Expected: 4")

	// Example 2: root = [3,3,null,4,2]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 2}
	// Expected output: 3
	println("Example 2 Output:", goodNodes(root2), "| Expected: 3")

	// Example 3: root = [1]
	root3 := &TreeNode{Val: 1}
	// Expected output: 1
	println("Example 3 Output:", goodNodes(root3), "| Expected: 1")

	// Additional Testcase
	// Input: root = [-1,5,-2,4,4,2,-2,null,null,-4,null,-2,3,null,-2,0,null,-1,null,-3,null,-4,-3,3,null,null,null,null,null,null,null,3,-3]
	// Expected output: 5
	root4 := &TreeNode{Val: -1}
	root4.Left = &TreeNode{Val: 5}
	root4.Right = &TreeNode{Val: -2}
	root4.Left.Left = &TreeNode{Val: 4}
	root4.Left.Right = &TreeNode{Val: 4}
	root4.Right.Left = &TreeNode{Val: 2}
	root4.Right.Right = &TreeNode{Val: -2}
	root4.Left.Left.Left = nil
	root4.Left.Left.Right = nil
	root4.Left.Right.Left = &TreeNode{Val: -4}
	root4.Left.Right.Right = nil
	root4.Right.Left.Left = &TreeNode{Val: -2}
	root4.Right.Left.Right = &TreeNode{Val: 3}
	root4.Right.Right.Left = nil
	root4.Right.Right.Right = &TreeNode{Val: -2}
	root4.Left.Right.Left.Left = &TreeNode{Val: 0}
	root4.Left.Right.Left.Right = nil
	root4.Right.Left.Right.Left = &TreeNode{Val: -1}
	root4.Right.Left.Right.Right = nil
	root4.Right.Right.Right.Left = &TreeNode{Val: -3}
	root4.Right.Right.Right.Right = nil
	root4.Left.Right.Left.Left.Left = &TreeNode{Val: -4}
	root4.Left.Right.Left.Left.Right = &TreeNode{Val: -3}
	root4.Right.Left.Right.Left.Left = &TreeNode{Val: 3}
	root4.Right.Left.Right.Left.Right = nil
	root4.Left.Right.Left.Left.Left.Left = nil
	root4.Left.Right.Left.Left.Left.Right = nil
	root4.Left.Right.Left.Left.Right.Left = nil
	root4.Left.Right.Left.Left.Right.Right = &TreeNode{Val: 3}
	root4.Right.Left.Right.Left.Left.Left = nil
	root4.Right.Left.Right.Left.Left.Right = &TreeNode{Val: -3}
	println("Additional Testcase Output:", goodNodes(root4), "| Expected: 5")

}
