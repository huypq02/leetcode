package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// If both p and q are less than the root value,
	// then the LCA must be in the left subtree
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// If both p and q are greater than the root value,
	// then the LCA must be in the right subtree
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	// If p and q are on different sides of the root,
	// or one of them is equal to the root value,
	// then the root is their lowest common ancestor
	return root
}

// Helper function to create a binary tree from a level-order array representation
// where -1 represents null values
func createBST(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}

	node := &TreeNode{Val: nums[index]}

	// Calculate the indices for left and right children
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	// Create left and right subtrees if within bounds
	if leftIndex < len(nums) {
		node.Left = createBST(nums, leftIndex)
	}
	if rightIndex < len(nums) {
		node.Right = createBST(nums, rightIndex)
	}

	return node
}

// Helper function to find a node with a given value in a binary tree
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	leftResult := findNode(root.Left, val)
	if leftResult != nil {
		return leftResult
	}

	return findNode(root.Right, val)
}

// Helper function to run a test case
func runTestCase(testNum int, root *TreeNode, p, q *TreeNode, expected int) {
	fmt.Printf("Example %d: ", testNum)
	result := lowestCommonAncestor(root, p, q)

	if result == nil {
		fmt.Printf("Failed - expected %d, got nil\n", expected)
	} else if result.Val != expected {
		fmt.Printf("Failed - expected %d, got %d\n", expected, result.Val)
	} else {
		fmt.Printf("Passed - result is %d\n", result.Val)
	}
}

// Helper function to print a simple visualization of the BST
func printBST(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	// Print current node
	if isLeft {
		fmt.Printf("%s├── %d\n", prefix, root.Val)
		prefix += "│   "
	} else {
		fmt.Printf("%s└── %d\n", prefix, root.Val)
		prefix += "    "
	}

	// Print children
	if root.Left != nil {
		printBST(root.Left, prefix, true)
	}
	if root.Right != nil {
		printBST(root.Right, prefix, false)
	}
}

func runExampleTest(exampleNum int, nums []int, pVal, qVal, expectedVal int) {
	fmt.Printf("\nExample %d:\n", exampleNum)

	// Create tree
	root := createBST(nums, 0)

	fmt.Println("Tree structure:")
	printBST(root, "", false)

	// Find target nodes
	p := findNode(root, pVal)
	q := findNode(root, qVal)

	// Run test case
	fmt.Printf("Finding LCA of nodes %d and %d (expected: %d):\n", pVal, qVal, expectedVal)
	runTestCase(exampleNum, root, p, q, expectedVal)
}

func main() {
	fmt.Println("Testing Lowest Common Ancestor of a Binary Search Tree")
	fmt.Println("-------------------------------------------------------")

	// Example 1: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8, expected = 6
	nums1 := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	runExampleTest(1, nums1, 2, 8, 6)

	// Example 2: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4, expected = 2
	// Reusing the same tree structure from Example 1
	runExampleTest(2, nums1, 2, 4, 2)

	// Example 3: root = [2,1], p = 2, q = 1, expected = 2
	nums3 := []int{2, 1}
	runExampleTest(3, nums3, 2, 1, 2)

	fmt.Println("\nNote: All test cases will show 'Failed' until the lowestCommonAncestor function is implemented.")
}
