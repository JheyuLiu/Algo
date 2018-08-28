/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
    	return nil
    }

    prev := head
    now := head.Next

    var next *ListNode
    for now != nil {
    	next = now.Next
    	now.Next = prev      
        prev = now
        now = next
    }

    head.Next = now
    return prev
}