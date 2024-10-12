package main

import (
    "fmt"
    "math"
)

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func calculateDepth(root *TreeNode, depth int) int{
    if root == nil {
        return depth-1
    }

    if root.Left == nil && root.Right == nil {
        return depth
    }
    left := calculateDepth(root.Left, depth+1)
    right := calculateDepth(root.Right, depth+1)
    fmt.Println(depth, ":", right)
    if int(math.Abs(float64(left - right))) > 1 {
        return -1
    }
    if (left > right) {
        return left
    }
    return right
}
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var depth int = 0
    if calculateDepth(root, depth+1) == -1  {
        return false
    }
    return true
}



func main(){
    // [3,9,20,null,null,15,7]

    tree := &TreeNode{
        Val: 3,
        Left: &TreeNode{9, nil, nil},
        Right: &TreeNode{
            Val: 20,
            Left: &TreeNode{15, nil, nil},
            Right: &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(isBalanced(tree))
    // [1,null,2,null,3]
    tree = &TreeNode{
        Val: 1,
        Left: nil,
        Right: &TreeNode{
            Val: 2,
            Left: nil,
            Right: &TreeNode{3, nil, nil},
        },
    }
    fmt.Println(isBalanced(tree))
}