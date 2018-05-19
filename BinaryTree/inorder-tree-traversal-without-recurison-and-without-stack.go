package main

import "fmt"

type tnode struct {
	data int
	left *tnode
	right *tnode
}

type btree struct {
	root *tnode
}

func (bt *btree) inOrder(root *tnode) {
	curr := root
	var prev *tnode
	for curr != nil {
		if curr.left == nil {
			fmt.Print(curr.data, " ")
			curr = curr.right
		} else {
			prev = curr.left
			for prev.right != nil && prev.right != curr{
				prev = prev.right
			}
				
			if prev.right == nil {
				prev.right = curr
				curr = curr.left
			} else {
				prev.right = nil
				fmt.Print(curr.data, " ")
				curr = curr.right
			}
		}
	}
}

func main() {
	bt := btree{root: &tnode{data: 1}}

	bt.root.left       = &tnode{data: 2}
	bt.root.right      = &tnode{data: 3} 
	bt.root.left.left  = &tnode{data: 4}
	bt.root.left.right = &tnode{data: 5}

	bt.inOrder(bt.root)
	fmt.Println()
}