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

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	root = searchBST(root, 5)
    fmt.Println(root)
}