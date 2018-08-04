/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, first *int, second *int) {
	if root == nil {
		return
	}
    if root.Val != (*first) && (*second) > root.Val {
    	(*second) = root.Val
    }

    help(root.Left, first, second)
    help(root.Right, first, second)
}

func findSecondMinimumValue(root *TreeNode) int {
	first := root.Val
	second := math.MaxInt64
	help(root, &first, &second) 
	if second == math.MaxInt64 {
		return -1
	}
	return second   
}