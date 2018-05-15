package main

import "fmt"

type node struct {
	data int
	left *node
	right *node
}

type bst struct {
	root *node
}

func (bst *bst) printPreorder(node *node) {
	if node == nil {
		return
	}

	fmt.Print(node.data, " ")

	bst.printPreorder(node.left)

	bst.printPreorder(node.right)
}

func (bst *bst) printInorder(node *node) {
	if node == nil {
		return
	}

	bst.printInorder(node.left)

	fmt.Print(node.data, " ")

	bst.printInorder(node.right)
}

func (bst *bst) printPostorder(node *node) {
	if node == nil {
		return
	}

	bst.printPostorder(node.left)

	bst.printPostorder(node.right)

	fmt.Print(node.data, " ")
}

func main() {
	BST := bst{root: &node{data: 1}}

	BST.root.left       = &node{data: 2}
	BST.root.right      = &node{data: 3} 
	BST.root.left.left  = &node{data: 4}
	BST.root.left.right = &node{data: 5}

	fmt.Println("Preorder traversal of binary tree is")
	BST.printPreorder(BST.root)

	fmt.Println("\nInorder traversal of binary tree is")
	BST.printInorder(BST.root)

	fmt.Println("\nPostorder traversal of binary tree is")
	BST.printPostorder(BST.root)
}