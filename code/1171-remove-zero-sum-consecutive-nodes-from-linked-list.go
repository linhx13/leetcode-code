package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := ListNode{Val: 0, Next: head}
	cur := dummy.Next
	m := make(map[int]*ListNode)
	sm := make(map[*ListNode]int)
	m[0] = &dummy
	s := 0
	for cur != nil {
		s += cur.Val
		if _, ok := m[s]; ok {
			for node := m[s].Next; node != cur; node = node.Next {
				delete(m, sm[node])
				delete(sm, node)
			}
			m[s].Next = cur.Next
		} else {
			m[s] = cur
			sm[cur] = s
		}
		cur = cur.Next
	}
	return dummy.Next
}
