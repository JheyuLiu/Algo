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

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    prev := head
    cur := head.Next

    for cur != nil {
        if prev.Val == cur.Val {
            prev.Next = cur.Next
        }else {
            prev = cur
        }
        cur = cur.Next
    }

    return head
}