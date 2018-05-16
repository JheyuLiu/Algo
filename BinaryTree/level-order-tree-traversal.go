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

func (bst *bst) getHight(node *node) int {
	if node == nil {
		return 0
	}

	left_hight  := bst.getHight(node.left) + 1
	right_hight := bst.getHight(node.right) + 1

	var max_hight int
	if left_hight >= right_hight {
		max_hight = left_hight
	} else {
		max_hight = right_hight
	}

	return max_hight
}

func (bst *bst) printGivenLevel(node *node, level int) {
	if node == nil {
		return
	}

	if level == 1 {
		fmt.Print(node.data, " ")
	} else if level > 1 {
		bst.printGivenLevel(node.left, level-1)
		bst.printGivenLevel(node.right, level-1)
	}
}

func (bst *bst) printLevelOrder(node *node) {
	height := bst.getHight(node)
	for i := 1; i <= height; i++ {
		bst.printGivenLevel(node, i)
	}
}

func main() {
	BST := bst{root: &node{data: 1}}

	BST.root.left       = &node{data: 2}
	BST.root.right      = &node{data: 3} 
	BST.root.left.left  = &node{data: 4}
	BST.root.left.right = &node{data: 5}

	fmt.Println("Level Order traversal of binary tree is")
	BST.printLevelOrder(BST.root)
}