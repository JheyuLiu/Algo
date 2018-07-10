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

func findTarget(root *TreeNode, k int) bool {
    var list []int
    inOrder(root, &list)

    if len(list) < 2 {
        return false
    }

    left := 0
    right := len(list)-1
    for ; left<right; {
        sum := list[left]+list[right]
        if sum > k {
            right = right-1
        }else if sum < k {
            left = left+1
        }else {
            return true
        }
    }

    return false
}

func inOrder(root *TreeNode, list *[]int) {
    if root == nil {
    	return
    }
    
    inOrder(root.Left, list)
    *list = append(*list, root.Val)
    inOrder(root.Right, list)
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(findTarget(root, 6))
}