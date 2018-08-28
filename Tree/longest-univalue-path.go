/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var path int

func help(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := help(root.Left)
    right := help(root.Right)
    
    var arrLeft, arrRight int
    if root.Left != nil && root.Val == root.Left.Val {
        arrLeft = arrLeft + left + 1
    }

    if root.Right != nil && root.Val == root.Right.Val {
        arrRight = arrRight + right + 1
    }

    if path > arrLeft + arrRight {
        path = path
    }else {
        path = arrRight + arrLeft
    }
    
    if arrLeft > arrRight {
        return arrLeft
    }
    return arrRight
}

func longestUnivaluePath(root *TreeNode) int {
    path = 0
    help(root)
    return path
}