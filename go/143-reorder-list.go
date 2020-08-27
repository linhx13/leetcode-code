package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	roots := make([]*ListNode, 2)
	i := 0
	node, last := head, (*ListNode)(nil)
	for ; i < (n+1)/2; i++ {
		last = node
		node = node.Next
	}
	last.Next = nil
	roots[0] = head
	for ; i < n; i++ {
		next := node.Next
		node.Next = roots[1]
		roots[1] = node
		node = next
	}
	n = 0
	node = roots[0]
	for roots[0] != nil || roots[1] != nil {
		if roots[n%2] != nil {
			roots[n%2] = roots[n%2].Next
		}
		node.Next = roots[(n+1)%2]
		node = node.Next
		n++
	}
}

func buildList(nums []int) *ListNode {
	var head, last *ListNode
	for _, n := range nums {
		node := &ListNode{Val: n}
		if last == nil {
			head = node
			last = node
		} else {
			last.Next = node
			last = node
		}
	}
	return head
}

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Print("->", head.Val)
	}
	fmt.Println()
}

func main() {
	// head := buildList([]int{1, 2, 3, 4})
	head := buildList([]int{1, 2, 3, 4, 5})
	printList(head)
	reorderList(head)
	printList(head)
}
