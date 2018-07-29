/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
    	return nil
    }

    if root.Val == val {
    	return root
    }

    lres := searchBST(root.Left, val)
    rres := searchBST(root.Right, val)

    if lres != nil {
    	return lres
    }

    if rres != nil {
    	return rres
    }

    return nil
}