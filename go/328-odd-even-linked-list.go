package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var evenHead, evenTail *ListNode
	cur := head
	for cur != nil {
		if cur.Next != nil {
			if evenTail != nil {
				evenTail.Next = cur.Next
			}
			evenTail = cur.Next
			if evenHead == nil {
				evenHead = evenTail
			}
			cur.Next = cur.Next.Next
		}
		if cur.Next == nil {
			break
		} else {
			cur = cur.Next
		}
	}
	cur.Next = evenHead
	if evenTail != nil {
		evenTail.Next = nil
	}
	return head
}

func print(head *ListNode) {
	res := []int{}
	for head != nil {
		fmt.Println(head.Val)
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	var head, cur *ListNode
	for _, n := range nums {
		node := &ListNode{Val: n}
		if cur == nil {
			cur = node
		} else {
			cur.Next = node
			cur = node
		}
		if head == nil {
			head = cur
		}
	}
	print(head)
	res := oddEvenList(head)
	print(res)
}
