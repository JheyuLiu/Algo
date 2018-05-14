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

func main() {
	BST := bst{root: &node{data: 1}}

	BST.root.left      = &node{data: 2}
	BST.root.right     = &node{data: 3} 
	BST.root.left.left = &node{data: 4}

	fmt.Println(BST.root.data)
	fmt.Println(BST.root.left.data)
	fmt.Println(BST.root.right.data)
	fmt.Println(BST.root.left.left.data)
}