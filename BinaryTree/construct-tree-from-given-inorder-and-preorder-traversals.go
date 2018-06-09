package main

import "fmt"

type node struct {
	data byte
	left *node
	right *node
}

type bt struct {
	root *node
}

var preIndex int

func search(in []byte, start int, end int, val byte) int {
	for i := start; i <= end; i++ {
		if in[i] == val {
			return i
		}
	}

	return -1
}

func (bt *bt) buildTree(in []byte, pre []byte, start int, end int) *node {
	if start > end {
		return nil
	}

	node := &node{data: pre[preIndex]}
	preIndex = preIndex + 1
	if start == end {
		return node
	}

	index := search(in, start, end, node.data)

	node.left = bt.buildTree(in, pre, start, index-1)
	node.right = bt.buildTree(in, pre, index+1, end)

	return node
}

func (bt *bt) printInorder(node *node) {
	if node == nil {
		return
	}

	bt.printInorder(node.left)

	fmt.Print(string(node.data), " ")

	bt.printInorder(node.right)
}

func main() {
	bt := bt{}
	in := []byte{'D', 'B', 'E', 'A', 'F', 'C'}
	pre := []byte{'A', 'B', 'D', 'E', 'C', 'F'}
	bt.root = bt.buildTree(in, pre, 0, len(in)-1)

	fmt.Println("Inorder traversal of constructed tree is : ")
	bt.printInorder(bt.root)
}