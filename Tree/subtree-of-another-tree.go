/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
	    return false
	}
	if s.Val != t.Val {
		return false
	}

    return help(s.Left, t.Left) && help(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
    }

	res := help(s, t)
	if !res {
		res = isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}

	return res
}