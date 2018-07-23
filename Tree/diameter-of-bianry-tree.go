/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, max *int) int {
	if root == nil {
		return -1
	}    

	left_val := help(root.Left, max)+1
	right_val := help(root.Right, max)+1
	if left_val + right_val > (*max) {
		(*max) = left_val + right_val
	}
	if left_val >= right_val {
		return left_val
	} 

	return right_val
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	help(root, &max)
	return max
}