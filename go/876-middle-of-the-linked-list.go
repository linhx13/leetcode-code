package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	l := 0
	node := head
	for node != nil {
		l++
		node = node.Next
	}
	node = head
	for i := 0; i < l/2; i++ {
		node = node.Next
	}
	return node
}
