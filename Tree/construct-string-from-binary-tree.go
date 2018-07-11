package main

import "fmt"
import "strconv"
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

func tree2str(t *TreeNode) string {
	var str string
	result := composeStr(t, &str)
	return result
}

func composeStr(node *TreeNode, str *string) string {
	if node == nil {
		return ""
	}
    (*str) = (*str) + strconv.Itoa(node.Val)

    if node.Left != nil || (node.Left == nil && node.Right != nil) {
    	(*str) = (*str) + "("
    	composeStr(node.Left, str)
    	(*str) = (*str) + ")"
    }
    if node.Right != nil {
    	(*str) = (*str) + "("
    	composeStr(node.Right, str)
    	(*str) = (*str) + ")"
    }

    return *str
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}

	fmt.Println(tree2str(root))
}