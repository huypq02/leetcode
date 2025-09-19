package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p != nil && q == nil) || (q != nil && p == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
		return false
	}
	if (p.Right == nil && q.Right != nil) || (p.Right != nil && q.Right == nil) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// Input: p = [1,2,3], q = [1,2,3]
	// Output: true
	p1, q1 :=
		&TreeNode{
			Val:   1,
			Left:  &TreeNode{2, nil, nil},
			Right: &TreeNode{3, nil, nil},
		},
		&TreeNode{
			Val:   1,
			Left:  &TreeNode{2, nil, nil},
			Right: &TreeNode{3, nil, nil},
		}
	println(isSameTree(p1, q1))

	// Input: p = [1,2], q = [1,null,2]
	// Output: false
	p2, q2 :=
		&TreeNode{
			Val:   1,
			Left:  &TreeNode{2, nil, nil},
			Right: nil,
		},
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{2, nil, nil},
		}
	println(isSameTree(p2, q2))

	// Input: p = [1,2,1], q = [1,1,2]
	// Output: false
	p3, q3 :=
		&TreeNode{
			Val:   1,
			Left:  &TreeNode{2, nil, nil},
			Right: &TreeNode{1, nil, nil},
		},
		&TreeNode{
			Val:   1,
			Left:  &TreeNode{1, nil, nil},
			Right: &TreeNode{2, nil, nil},
		}
	println(isSameTree(p3, q3))

	// Input: p = [], q = [0]
	// Output: false
	p4, q4 :=
		(*TreeNode)(nil),
		&TreeNode{0, nil, nil}
	println(isSameTree(p4, q4))
}
