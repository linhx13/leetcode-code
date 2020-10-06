package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := ListNode{}
	p := &dummy
	for head != nil {
		c := 0
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
			c++
		}
		if c == 0 {
			p.Next = head
			p = p.Next
		} else {
			p.Next = nil
		}
		head = head.Next
	}
	return dummy.Next
}
