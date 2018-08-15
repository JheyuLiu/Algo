/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res bool

func help(root *TreeNode) int {
    if root == nil {
        return 0
    }

    ldepth := help(root.Left)
    rdepth := help(root.Right)
    
    if ldepth >= rdepth {
        if ldepth-rdepth > 1 {
            res = false
        }
        return ldepth+1
    }
    if rdepth - ldepth > 1 {
        res = false
    }
    return rdepth+1
}

func isBalanced(root *TreeNode) bool {
    res = true
    help(root)
    return res
}