package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	ll := 0
	for p := root; p != nil; p = p.Next {
		ll++
	}
	n := ll / k
	d := ll - n*k
	res := []*ListNode{}
	for p := root; p != nil; {
		cur := p
		var last *ListNode
		for c := 0; c < n; c++ {
			last, p = p, p.Next
		}
		if d != 0 {
			last, p = p, p.Next
			d--
		}
		last.Next = nil
		res = append(res, cur)
	}
	for len(res) < k {
		res = append(res, nil)
	}
	return res
}
