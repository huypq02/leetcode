package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left, maxDiameter)
	right := depth(root.Right, maxDiameter)
	if *maxDiameter < left+right {
		*maxDiameter = left + right
	}

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	// Reset before calculation
	maxDiameter := 0

	// Pass the address of maxDiameter to calculate
	// the max sum of left and right subtree depths for ANY NODE, NOT just the ROOT
	depth(root, &maxDiameter)

	return maxDiameter
}

func main() {

	// Input: root = [1,2,3,4,5]
	// Output: 3
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{4, nil, nil},
			Right: &TreeNode{5, nil, nil},
		},
		Right: &TreeNode{3, nil, nil},
	}
	println(diameterOfBinaryTree(root1))

	// Input: root = [1,2]
	// Output: 1
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: nil,
	}
	println(diameterOfBinaryTree(root2))

	// Input: root = [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2]
	// Output: 8
	root3 := &TreeNode{
		Val:  4,
		Left: &TreeNode{-7, nil, nil},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -9,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: &TreeNode{-1, nil, nil},
						},
						Right: &TreeNode{
							Val:   6,
							Left:  &TreeNode{-4, nil, nil},
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: -7,
					Left: &TreeNode{
						Val:   -6,
						Left:  &TreeNode{5, nil, nil},
						Right: nil,
					},
					Right: &TreeNode{
						Val: -6,
						Left: &TreeNode{
							Val:   9,
							Left:  &TreeNode{-2, nil, nil},
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				Val:   -3,
				Left:  &TreeNode{-4, nil, nil},
				Right: nil,
			},
		},
	}
	println(diameterOfBinaryTree(root3))

}
