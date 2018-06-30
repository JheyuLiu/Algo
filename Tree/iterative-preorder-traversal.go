package main

import "fmt"

type tNode struct {
	data int
	left *tNode
	right *tNode
}

type btree struct {
	root *tNode
}

type sNode struct {
	node *tNode
	next *sNode
}

type stack struct {
	top *sNode
}

func (st *stack) push(node *tNode) {
	new_snode := &sNode{node: node, next: nil}
	new_snode.next = st.top.next
	st.top.next = new_snode
}

func (st *stack) pop() *tNode {
	fmt.Print(st.top.next.node.data, " ")
	node := st.top.next.node
	st.top.next = st.top.next.next

	return node
}

func (bt *btree) iterativePreorder(node *tNode) {
	if node == nil {
		return
	}

	snode := &sNode{node: node, next: nil}
	stk := &stack{top: &sNode{node: nil, next: snode}}
	
	for stk.top.next != nil {
		tnode := stk.pop()
		if tnode.right != nil {
			stk.push(tnode.right)
		}
		if tnode.left != nil {
			stk.push(tnode.left)
		}
	}
}

func main() {
	bt := btree{root: &tNode{data: 10}}

	bt.root.left       = &tNode{data: 8}
	bt.root.right      = &tNode{data: 2} 
	bt.root.left.left  = &tNode{data: 3}
	bt.root.left.right = &tNode{data: 5}
	bt.root.right.left = &tNode{data: 2}

	bt.iterativePreorder(bt.root)
	fmt.Println()
}