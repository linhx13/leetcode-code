package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	var res *ListNode
	for next := (*ListNode)(nil); head != nil; head = next {
		next = head.Next
		node := res
		last := (*ListNode)(nil)
		for node != nil && node.Val < head.Val {
			last = node
			node = node.Next
		}
		if last == nil {
			head.Next = res
			res = head
		} else {
			head.Next = last.Next
			last.Next = head
		}
	}
	return res
}
