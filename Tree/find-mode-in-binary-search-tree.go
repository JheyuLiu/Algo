/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var max int

func help(root *TreeNode, m *map[int]int) {
    if root == nil {
        return
    }
    (*m)[root.Val] = (*m)[root.Val] + 1
    if (*m)[root.Val] > max {
        max = (*m)[root.Val]
    }
    help(root.Left, m)
    help(root.Right, m)
}

func findMode(root *TreeNode) []int {
    m := map[int]int{}
    max = 0
    help(root, &m)
    
    var s []int
    for k, _ := range m {
        if max == m[k] {
            s = append(s, k)
        }
    }
    return s
}