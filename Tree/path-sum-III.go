/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, sum int, pathCount int) int {
    if root == nil {
        return 0
    }

    sum = sum - root.Val
    if sum == 0 {
        pathCount = pathCount + 1
    }
    pathCount = pathCount + help(root.Left, sum, 0) + help(root.Right, sum, 0)
    return pathCount
}

func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }

    total := help(root, sum, 0)
    return total + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}