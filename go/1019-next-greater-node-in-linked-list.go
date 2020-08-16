package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	stack := [][]int{}
	res := []int{}
	cnt := 0
	for head != nil {
		res = append(res, 0)
		for len(stack) > 0 && stack[len(stack)-1][0] < head.Val {
			t := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			res[t[1]] = head.Val
		}
		stack = append(stack, []int{head.Val, cnt})
		cnt++
		head = head.Next
	}
	return res
}
