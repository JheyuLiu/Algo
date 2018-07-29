/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func PathsRecur(str string, res []string, root *TreeNode) []string {
	if root == nil {
    	return nil
    }

    if str == "" {
    	str = str + strconv.Itoa(root.Val)
    } else {
    	str = str + "->" + strconv.Itoa(root.Val)
    }

    if root.Left == nil && root.Right == nil {
    	res = append(res, str)
    }

    if root.Left != nil {
    	res = PathsRecur(str, res, root.Left)
    } 
    if root.Right != nil {
    	res = PathsRecur(str, res, root.Right)
    }

    return res
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0, 16)
	var str string
    res = PathsRecur(str, res, root)
    
    return res
}