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

func (bt *bt) getHight(node *node) int {
	if node == nil {
		return 0
	}

	left_hight  := bt.getHight(node.left) + 1
	right_hight := bt.getHight(node.right) + 1

	var max_hight int
	if left_hight >= right_hight {
		max_hight = left_hight
	} else {
		max_hight = right_hight
	}

	return max_hight
}

func (bt *bt) printGivenLevel(node *node, level int) {
	if node == nil {
		return
	}

	if level == 1 {
        fmt.Print(node.data, " ")	
    } else {
		bt.printGivenLevel(node.left, level-1)
		bt.printGivenLevel(node.right, level-1)
	}
}

func (bt *bt) printLevelOrder(node *node) {
	height := bt.getHight(bt.root)
	for i := 1; i <= height; i++ {
		bt.printGivenLevel(bt.root, i)
		fmt.Println()
	}
}

func main() {
	BT := bt{root: &node{data: 1}}

	BT.root.left       = &node{data: 2}
	BT.root.right      = &node{data: 3} 
	BT.root.left.left  = &node{data: 4}
	BT.root.left.right = &node{data: 5}
	BT.root.right.right = &node{data: 6}

	BT.printLevelOrder(BT.root)
}