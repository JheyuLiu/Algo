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

func main() {
	BT := bt{root: &node{data: 1}}

	BT.root.left      = &node{data: 2}
	BT.root.right     = &node{data: 3} 
	BT.root.left.left = &node{data: 4}

	fmt.Println(BT.root.data)
	fmt.Println(BT.root.left.data)
	fmt.Println(BT.root.right.data)
	fmt.Println(BT.root.left.left.data)
}