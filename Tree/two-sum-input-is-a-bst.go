/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func find(node *TreeNode, k int, tmap *map[int]int) bool {
    if node == nil {
        return false
    }

    num := k-node.Val
    _, ok := (*tmap)[num]
    if ok {
        return true
    } else {
        (*tmap)[node.Val] = 0
    }

    res_left := find(node.Left, k, tmap)
    res_right := find(node.Right,k, tmap)

    return res_left || res_right
}

func findTarget(root *TreeNode, k int) bool {
    tmap := make(map[int]int)
    return find(root, k, &tmap)
}
