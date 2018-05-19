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

func (st *stack) pop(node *tNode) {
	fmt.Print(node.data, " ")
    st.top.next = st.top.next.next
}

func (bt *btree) inOrder(node *tNode) {
	snode := &sNode{node: node, next: nil}
	stk := &stack{top: &sNode{node: nil, next: snode}}
	
	curr := node.left
	done := false
	for !done {
		if curr != nil {
			stk.push(curr)
			curr = curr.left
		} else {
			if stk.top.next != nil {
				curr = stk.top.next.node.right
				stk.pop(stk.top.next.node)
			} else {
				done = true
			}
		}
	}
}

func main() {
	bt := btree{root: &tNode{data: 1}}

	bt.root.left       = &tNode{data: 2}
	bt.root.right      = &tNode{data: 3} 
	bt.root.left.left  = &tNode{data: 4}
	bt.root.left.right = &tNode{data: 5}

	bt.inOrder(bt.root)
	fmt.Println()
}