/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    var sum int
    return convertHelp(root, &sum)
}

func convertHelp(root *TreeNode, sum *int) *TreeNode {
    if root == nil {
        return nil
    }

    convertHelp(root.Right, sum)
    root.Val = root.Val + (*sum)
    (*sum) = root.Val
    convertHelp(root.Left, sum)

    return root
}