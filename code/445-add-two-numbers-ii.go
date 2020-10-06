package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := getLen(l1)
	len2 := getLen(l2)
	res, carry := helper(l1, l2, len1, len2)
	fmt.Println(res, carry)
	if carry > 0 {
		res = &ListNode{Val: carry, Next: res}
	}
	return res
}

func getLen(head *ListNode) int {
	n := 0
	for ; head != nil; head = head.Next {
		n++
	}
	return n
}

func helper(l1 *ListNode, l2 *ListNode, len1 int, len2 int) (*ListNode, int) {
	if len1 == 1 && len2 == 1 {
		val := l1.Val + l2.Val
		return &ListNode{Val: val % 10}, val / 10
	}
	res := (*ListNode)(nil)
	carry := 0
	if len1 > len2 {
		res, carry = helper(l1.Next, l2, len1-1, len2)
		val := l1.Val + carry
		res = &ListNode{Val: val % 10, Next: res}
		carry = val / 10
	} else if len2 > len1 {
		res, carry = helper(l1, l2.Next, len1, len2-1)
		val := l2.Val + carry
		res = &ListNode{Val: val % 10, Next: res}
		carry = val / 10
	} else {
		res, carry = helper(l1.Next, l2.Next, len1-1, len2-1)
		val := l1.Val + l2.Val + carry
		res = &ListNode{Val: val % 10, Next: res}
		carry = val / 10
	}
	return res, carry
}
