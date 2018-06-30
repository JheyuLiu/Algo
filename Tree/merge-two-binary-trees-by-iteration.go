package main

import (
    "fmt"
    "github.com/golang-collections/collections/stack"
)

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    stk := stack.New()
    stk.Push([2]*TreeNode{t1, t2})

	for stk.Peek() != nil {
		t := stk.Pop()
		t1_node := t.([2]*TreeNode)[0]
		t2_node := t.([2]*TreeNode)[1]

		if (t1_node == nil || t2_node == nil) {
            continue
		}

        t1_node.Val = t1_node.Val + t2_node.Val 
		if t1_node.Left == nil {
            t1_node.Left = t2_node.Left
		} else {
			stk.Push([2]*TreeNode{t1_node.Left, t2_node.Left})
		}

		if t1_node.Right == nil {
			t1_node.Right = t2_node.Right
		} else {
			stk.Push([2]*TreeNode{t1_node.Right, t2_node.Right})
		}
	}

	return t1
}