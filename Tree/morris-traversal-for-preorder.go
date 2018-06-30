package main

import "fmt"

type node struct {
	data int
	left *node
	right *node
}

type btree struct {
	root *node
}

func (bt *btree) morrisTraversalPreorder(root *node) {
	curr := root
	var prev *node
	for curr != nil {
		if curr.left == nil {
			fmt.Print(curr.data, " ")
			curr = curr.right
		} else {
			prev = curr.left
			for prev.right != nil && prev.right != curr {
				prev = prev.right
			}

			if prev.right != nil {
				prev.right = nil
				curr = curr.right
			} else {
				prev.right = curr
				fmt.Print(curr.data, " ")
				curr = curr.left
			}
		}
	}
}

func main() {
	bt := btree{root: &node{data: 1}}

	bt.root.left             = &node{data: 2}
	bt.root.right            = &node{data: 3} 
	bt.root.left.left        = &node{data: 4}
	bt.root.left.right       = &node{data: 5}
	bt.root.right.left       = &node{data: 6}
	bt.root.right.right      = &node{data: 7}
	bt.root.left.left.left   = &node{data: 8}
	bt.root.left.left.right  = &node{data: 9}
	bt.root.left.right.left  = &node{data: 10}
	bt.root.left.right.right = &node{data: 11}

	bt.morrisTraversalPreorder(bt.root)
	fmt.Println()
}