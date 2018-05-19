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

func (bt *bt) printPreorder(node *node) {
	if node == nil {
		return
	}

	fmt.Print(node.data, " ")

	bt.printPreorder(node.left)

	bt.printPreorder(node.right)
}

func (bt *bt) printInorder(node *node) {
	if node == nil {
		return
	}

	bt.printInorder(node.left)

	fmt.Print(node.data, " ")

	bt.printInorder(node.right)
}

func (bt *bt) printPostorder(node *node) {
	if node == nil {
		return
	}

	bt.printPostorder(node.left)

	bt.printPostorder(node.right)

	fmt.Print(node.data, " ")
}

func main() {
	BT := bt{root: &node{data: 1}}

	BT.root.left       = &node{data: 2}
	BT.root.right      = &node{data: 3} 
	BT.root.left.left  = &node{data: 4}
	BT.root.left.right = &node{data: 5}

	fmt.Println("Preorder traversal of binary tree is")
	BT.printPreorder(BT.root)

	fmt.Println("\nInorder traversal of binary tree is")
	BT.printInorder(BT.root)

	fmt.Println("\nPostorder traversal of binary tree is")
	BT.printPostorder(BT.root)
}