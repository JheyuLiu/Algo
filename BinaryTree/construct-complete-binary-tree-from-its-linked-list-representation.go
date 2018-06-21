package main

import "fmt"

type queue struct {
	qdata []*node
	start int
	last int
}

type ListNode struct {
	data int
	next *ListNode
}

type List struct {
	head *ListNode
}

type node struct {
	data int
	left *node
	right *node
}

type bt struct {
	root *node
}

func (q *queue) push(node *node) {
	q.qdata[q.last] = node
	q.last++
}

func (q *queue) pop() *node {
	if q.start == q.last {
		q.start = 0
		q.last = 0
	}

	val := q.qdata[q.start]
	q.start++

	return val
}

func (list *List) length() int {
	var length int
	node := list.head
	
	for node != nil {
		node = node.next
		length = length + 1
	}

	return length
}

func (list *List) push(val int) {
	node := &ListNode{data: val, next: nil}
	if(list.head == nil) {
		list.head = node
		return
	}

	tmp_node := list.head
	for tmp_node.next != nil {
		tmp_node = tmp_node.next
	}

	tmp_node.next = node
}

func (list *List) converList2Binary(bt *bt) {
    size := list.length()
    q := queue{}
    q.qdata = make([]*node, size)

    bt.root = &node{data: list.head.data}
    q.push(bt.root)

    head := list.head.next
    for head != nil {
    	parent := q.pop()
    	var leftchild, rightchild *node

    	leftchild = &node{data: head.data}
    	q.push(leftchild)

    	head = head.next
    	if head != nil {
    		rightchild = &node{data: head.data}
    		q.push(rightchild)
    		head = head.next
    	}

    	parent.left = leftchild
    	parent.right = rightchild
    }
}


func (bt *bt) inorderTraversal(node *node) {
	if node != nil {
		bt.inorderTraversal(node.left)
		fmt.Print(node.data, " ")
		bt.inorderTraversal(node.right)
	}
}

func main() {
	list := List{}
    list.push(10);
    list.push(12); 
    list.push(15);
    list.push(25);
    list.push(30);
    list.push(36);

    bt := bt{}
    list.converList2Binary(&bt)
    
    fmt.Println("Inorder Traversal of the constructed Binary Tree is:")
    bt.inorderTraversal(bt.root)
    fmt.Println()
}