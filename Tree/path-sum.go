/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }

    sum = sum - root.Val
    if sum == 0 && root.Left == nil && root.Right == nil {
        return true
    }
    return help(root.Left, sum) || help(root.Right, sum)
}

func hasPathSum(root *TreeNode, sum int) bool {
    return help(root, sum)
}