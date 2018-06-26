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

func isStructSame(anode *node, bnode *node) bool {
	if anode == nil && bnode == nil {
		return true;
	}

	if anode != nil && bnode != nil && isStructSame(anode.left, bnode.left) && isStructSame(anode.right, bnode.right) {
		return true
	}

	return false
}

func (bt *bt) isFoldable(node *node) bool {
	if node == nil {
		return true
	}

	bt.mirror(node.left)
	res := isStructSame(node.left, node.right)
	bt.mirror(bt.root)

	return res
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

func main() {
	bt := bt{}
	bt.root = &node{data: 10}
	bt.root.left = &node{data: 7}
	bt.root.right = &node{data: 15}
	bt.root.left.left = &node{data: 5}
	bt.root.right.left = &node{data: 11}

	if bt.isFoldable(bt.root) == true {
		fmt.Print("tree is foldable")
	} else {
		fmt.Print("tree is not foldable")
	}
	fmt.Println()
}