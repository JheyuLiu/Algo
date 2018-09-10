/**
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
 * linkedList.get(1);            // returns 2
 * linkedList.deleteAtIndex(1);  // now the linked list is 1->3
 * linkedList.get(1);            // returns 3
 */
type MyLinkedList struct {
    head, tail *node
    size int
}

type node struct {
    val int
    next *node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    t := &node{}
    h := &node{next: t}

    return MyLinkedList{
        head: h,
        tail: t,
    }
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (list *MyLinkedList) AddAtHead(val int) {
    n := &node{val: val}
    n.next = list.head.next
    list.head.next = n

    list.size++
}

/** Append a node of value val to the last element of the linked list. */
func (list *MyLinkedList) AddAtTail(val int) {
    list.tail.val = val
    n := &node{}
    list.tail.next = n
    list.tail = n

    list.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (list *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > list.size {
        return
    } else if index == 0 {
        list.AddAtHead(val)
        return
    } else if index == list.size {
        list.AddAtTail(val)
        return
    }

    i, cur := -1, list.head
    for i+1 < index {
        i++
        cur = cur.next
    }

    n := &node{val: val}
    n.next = cur.next
    cur.next = n

    list.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (list *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= list.size {
        return
    }

    i, cur := -1, list.head
    for i+1 < index {
        i++
        cur = cur.next
    }

    cur.next = cur.next.next
    list.size--
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (list *MyLinkedList) Get(index int) int {
    if index < 0 || index >= list.size {
        return -1
    }

    i, cur := 0, list.head.next
    for i < index {
        i++
        cur = cur.next
    }

    return cur.val
}