package main

import "fmt"

const SIZE = 100

type node struct {
	data int
	left *node
	right *node
}

type btree struct {
	root *node
}

type stack struct {
	items [SIZE]*node
	top int
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
func (stk *stack) push(node *node) {
	stk.top = stk.top + 1
	stk.items[stk.top] = node
}

func (stk *stack) pop() *node {
	if stk.top < 0 {
		return nil
	}
	node := stk.items[stk.top]
	stk.top = stk.top - 1
	return node
}

func (stk *stack) peek() *node {
	if stk.top < 0 {
		return nil
	}
	return stk.items[stk.top]
}

func (bt *btree) postOrderIterative(node *node) {
	s := stack{top: -1}
    if node == nil {
		return;
	}

	for  {
		for node != nil {
			if node.right != nil {
				s.push(node.right)
			}
			s.push(node)
			node = node.left
		}
		
		pop_node := s.pop()
		if pop_node != nil && pop_node.right == s.peek() {
			node = s.pop()
			s.push(pop_node)
		} else {
			fmt.Print(pop_node.data, " ")
			node = nil
		}


		if s.top < 0 {
			return;
		}
	}
}

func main() {
	bt := btree{root: &node{data: 1}}

	bt.root.left        = &node{data: 2}
	bt.root.right       = &node{data: 3} 
	bt.root.left.left   = &node{data: 4}
	bt.root.left.right  = &node{data: 5}
	bt.root.right.left  = &node{data: 6}
	bt.root.right.right = &node{data: 7}

	bt.postOrderIterative(bt.root)
	fmt.Println()
}