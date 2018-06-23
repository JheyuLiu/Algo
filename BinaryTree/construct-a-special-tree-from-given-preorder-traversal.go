package main

import "fmt"

type node struct {
	data int
	left *node
	right *node
}

type bt struct {
	root *node
	index int
}

func (bt *bt) constructTreeUT(pre [5]int, preLN [5]byte, index int, tmp_node *node) (*node) {
	tmp_node = &node{data: pre[bt.index]}
	bt.index = index + 1
	if preLN[index] == 'N' {
		tmp_node.left = bt.constructTreeUT(pre, preLN, bt.index, tmp_node.left)
		tmp_node.right = bt.constructTreeUT(pre, preLN, bt.index, tmp_node.right)
	}

	return tmp_node
}

func (bt *bt) constructTree(pre [5]int, preLN [5]byte, n int) (node *node) {
    return bt.constructTreeUT(pre, preLN, bt.index, bt.root)
}

func (bt *bt) printInorder(node *node) {
	if node == nil {
		return
	}

	bt.printInorder(node.left)
	fmt.Print(node.data, " ")
	bt.printInorder(node.right)
}

func main() {
	pre := [5]int{10, 30, 20, 5, 15}
	preLN := [5]byte{'N', 'N', 'L', 'L', 'L'}
	n := 5

	bt := bt{}
	bt.root = bt.constructTree(pre, preLN, n)

	fmt.Println("Following is Inorder Traversal of the Constructed Binary Tree: ")
	bt.printInorder(bt.root)
	fmt.Println()
}