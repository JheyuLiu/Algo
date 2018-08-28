package main

import "fmt"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    /*var count = 1
    tmp := head
    for tmp.Next != nil {
    	count = count + 1
    	tmp = tmp.Next
    }

    var index int
    for head.Next != nil {
        if index == count/2 {
        	break;
        }
        index = index + 1
        head = head.Next
    }*/
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
    head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	res_node := middleNode(head)
	fmt.Println("Node",res_node.Val,"from this list")
}