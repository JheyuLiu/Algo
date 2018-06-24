package main

import "fmt"

type node struct {
	data int
	left *node
	right *node
}

type bt struct {
	root *node
}

func (bt *bt) mirror(node *node) {
	if node == nil {
		return
	}

	bt.mirror(node.left)
	bt.mirror(node.right)

	tmp_node := node.left
	node.left = node.right
	node.right = tmp_node
}

func (bt *bt) inOrder(node *node) {
	if node == nil {
		return
	}

	bt.inOrder(node.left)
	fmt.Print(node.data, " ")
	bt.inOrder(node.right)
}

func main() {
	bt := bt{}
	bt.root = &node{data: 1}
	bt.root.left = &node{data: 2}
	bt.root.right = &node{data: 3}
	bt.root.left.left = &node{data: 4}
	bt.root.left.right = &node{data: 5}

	fmt.Println("Inorder traversal of the constructed tree is")
	bt.inOrder(bt.root)
	fmt.Println()

	bt.mirror(bt.root)

	fmt.Println("Inorder traversal of the mirror tree is")
	bt.inOrder(bt.root)
	fmt.Println()
}