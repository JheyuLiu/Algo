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

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    rever_node := reverseList(slow)
    return compare(head, rever_node)
}

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

func compare(f *ListNode, s *ListNode) bool {
    for f != nil && s != nil {
        if f.Val != s.Val {
            return false
        }
        f = f.Next
        s = s.Next
    }
    return true
}