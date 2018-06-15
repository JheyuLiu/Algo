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

func search(in []int, val int) int {
	for i := 0; i < len(in); i++ {
		if in[i] == val {
			return i
		}
	}
	return -1
}

func (bt *bt) buildTree(in []int, level []int) *node {

    if len(in) < 1 {
		return nil
	}
	n := &node{}
	var index int
	
		for i := 0; i < len(level); i++ {
			p := search(in, level[i])
			if p != -1 {
				n = &node{data: in[p]}
				index = p
				break;
			}
		}

	n.left = bt.buildTree(in[0:index], level)
	n.right = bt.buildTree(in[index+1:len(in)], level)

	return n
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
	bt := bt{}
	in := []int{4, 8, 10, 12, 14, 20, 22}
	level := []int{20, 8, 22, 4, 12, 10, 14}
	bt.root = bt.buildTree(in, level)

	fmt.Println("Inorder traversal of the constructed tree is")
	bt.printInorder(bt.root)
}