/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
    	return 0
    }

    left_count := maxDepth(root.Left) + 1
    right_count := maxDepth(root.Right) + 1

    if left_count >= right_count {
    	return left_count
    }
    return right_count
}