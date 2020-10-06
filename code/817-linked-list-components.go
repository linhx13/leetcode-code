package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	res := 0
	m := make(map[int]bool)
	for _, n := range G {
		m[n] = true
	}

	for head != nil {
		if m[head.Val] && (head.Next == nil || !m[head.Next.Val]) {
			res++
		}
		head = head.Next
	}
	return res
}
