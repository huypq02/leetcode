package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	var buildNode func(preorder []int, inorder []int) *TreeNode
	buildNode = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {
			return nil
		}

		// Identify the rootVal of the tree
		rootVal := preorder[0]
		// Create root node
		treenode := &TreeNode{Val: rootVal}
		// Find the current root index
		var rootIdx int
		for i, v := range inorder {
			if v == rootVal {
				rootIdx = i
			}
		}

		// The left subtree of the current tree
		treenode.Left = buildNode(preorder[1:1+rootIdx], inorder[:rootIdx])
		// The right subtree of the current tree
		treenode.Right = buildNode(preorder[1+rootIdx:], inorder[rootIdx+1:])

		return treenode
	}

	return buildNode(preorder, inorder)
}

func main() {
	// Example 1
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	fmt.Println("Testcase 1:")
	fmt.Println("preorder:", intSliceToString(preorder1))
	fmt.Println("inorder:", intSliceToString(inorder1))
	fmt.Println("Expected Output: [3,9,20,null,null,15,7]")
	tree1 := buildTree(preorder1, inorder1)
	fmt.Println("buildTree Output:", serializeTree(tree1))

	// Example 2
	preorder2 := []int{-1}
	inorder2 := []int{-1}
	fmt.Println("\nTestcase 2:")
	fmt.Println("preorder:", intSliceToString(preorder2))
	fmt.Println("inorder:", intSliceToString(inorder2))
	fmt.Println("Expected Output: [-1]")
	tree2 := buildTree(preorder2, inorder2)
	fmt.Println("buildTree Output:", serializeTree(tree2))
}

func intSliceToString(nums []int) string {
	s := "["
	for i, v := range nums {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%d", v)
	}
	s += "]"
	return s
}

// Helper to serialize tree as level-order (LeetCode format)
func serializeTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	res := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, "null")
		}
	}
	// Remove trailing "null"s
	i := len(res) - 1
	for i >= 0 && res[i] == "null" {
		i--
	}
	return "[" + fmt.Sprintf("%s", joinStrings(res[:i+1], ",")) + "]"
}

func joinStrings(arr []string, sep string) string {
	if len(arr) == 0 {
		return ""
	}
	s := arr[0]
	for _, v := range arr[1:] {
		s += sep + v
	}
	return s
}
