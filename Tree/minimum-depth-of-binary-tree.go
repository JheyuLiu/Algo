/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }else if root.Left == nil && root.Right != nil {
        return minDepth(root.Right) + 1
    }else if root.Left != nil && root.Right == nil {
        return minDepth(root.Left) + 1
    }

    left := minDepth(root.Left) + 1
    right := minDepth(root.Right) + 1
    if left < right {
        return left
    }
    return right
}