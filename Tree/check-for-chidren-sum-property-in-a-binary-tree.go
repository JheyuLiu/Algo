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

func isSumProperty(node *node) bool {
	if node == nil {
		return true
	}

	if node.left == nil && node.right == nil {
		return true
	}

	var left_sum, right_sum int
	if node.left != nil {
		left_sum = node.left.data
	}
	if node.right != nil {
		right_sum = node.right.data
	}

	sum := left_sum + right_sum
	if sum == node.data && isSumProperty(node.left) && isSumProperty(node.right){
		return true
	}

	return false
}

func main() {
	bt := bt{}
	bt.root = &node{data: 10}
	bt.root.left = &node{data: 8}
	bt.root.right = &node{data: 2}
	bt.root.left.left = &node{data: 3}
	bt.root.left.right = &node{data: 5}
	bt.root.right = &node{data: 2}

	if isSumProperty(bt.root) {
		fmt.Println("The given tree satisfies the children sum property")
	} else {
		fmt.Println("The given tree does not satisfy the children sum property")
	}
}