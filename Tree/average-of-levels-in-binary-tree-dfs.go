package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    var count []int

    average(root, &res, 0, &count)
    for i := 0; i < len(res); i++ {
        res[i] = res[i] / float64(count[i])
    }

    return res
}

func average(node *TreeNode, res *[]float64, i int, count *[]int) {
    if node == nil {
        return
    }

    if i < len(*res) {
        (*res)[i] = (*res)[i] + float64(node.Val)
        (*count)[i] = (*count)[i] + 1
    } else {
        *res = append(*res, 1)
        (*res)[i] = float64(node.Val)
        *count = append(*count, 1)
        (*count)[i] = 1
    }

    average(node.Left, res, i+1, count)
    average(node.Right, res, i+1, count)
}

func main() {
    root           := &TreeNode{Val: 3}
    root.Left       = &TreeNode{Val: 9}
    root.Right      = &TreeNode{Val: 20}
    root.Right.Left  = &TreeNode{Val: 15}
    root.Right.Right = &TreeNode{Val: 7}

    res := averageOfLevels(root)

    fmt.Println(res)
}