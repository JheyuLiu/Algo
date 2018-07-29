/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result [][]int

func levelOrderTraversal(root *TreeNode, level int) *TreeNode {
    if root == nil {
    	return nil
    }
    
    levelOrderTraversal(root.Left, level+1)
    levelOrderTraversal(root.Right, level+1)

    result[level] = append(result[level], root.Val)
    return root
}

func levelOrderBottom(root *TreeNode) [][]int {
	result = make([][]int, 100)
    levelOrderTraversal(root, 0)
    res := [][]int{}

	for i := len(result)-1; i >= 0; i-- {
		if result[i] != nil {
			res = append(res, result[i])
		}
	}
    return res
}