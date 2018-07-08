/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Left == nil && root.Right == nil) {
    	return root
    }

    var left, right *TreeNode
    if root.Left != nil {
    	left = invertTree(root.Left)
    }
    if root.Right != nil {
    	right = invertTree(root.Right)
    }

    tmp := left
    root.Left = right
    root.Right = tmp

    return root
}