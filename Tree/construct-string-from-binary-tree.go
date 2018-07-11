package main

import "fmt"
import "strconv"
import "bytes"
import "github.com/golang-collections/collections/stack"
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
	if t == nil {
		return ""
	}
	stk := stack.New()
	vmap := make(map[*TreeNode]int)
    var buffer bytes.Buffer

    stk.Push(t)
    for stk.Len() > 0 {
        node := stk.Peek()
        _, visited := vmap[node.(*TreeNode)]
        if visited {
        	stk.Pop()
        	buffer.WriteString(")")
        }else {
        	vmap[node.(*TreeNode)] = 0
        	buffer.WriteString("(")
            buffer.WriteString(strconv.Itoa(node.(*TreeNode).Val))

        	if node.(*TreeNode).Left == nil && node.(*TreeNode).Right != nil {
        		buffer.WriteString("()")
        	}
        	if node.(*TreeNode).Right != nil {
        		stk.Push(node.(*TreeNode).Right)
        	}
        	if node.(*TreeNode).Left != nil {
        		stk.Push(node.(*TreeNode).Left)
            }
        }
    }

    return buffer.String()
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(tree2str(root))
}