/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(x int) int {
	y := (x >> 31)
	return (x ^ y) - y
}

func help(root *TreeNode, tilt *int) int {
    if root == nil {
    	return 0
    }

    left := help(root.Left, tilt)
    right := help(root.Right, tilt)

    (*tilt) = (*tilt) + abs(left - right)
    return left + right + root.Val
}

func findTilt(root *TreeNode) int {
	var tilt int
	help(root, &tilt)
    return tilt
}