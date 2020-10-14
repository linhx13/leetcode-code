package main

import "fmt"

func maxWidthRamp(A []int) int {
	stack := []int{}
	n := len(A)
	res := 0
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		if A[stack[len(stack)-1]] > A[i] {
			stack = append(stack, i)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if len(stack) != 0 && stack[len(stack)-1] == i {
			stack = stack[0 : len(stack)-1]
		}
		for len(stack) != 0 && A[stack[len(stack)-1]] <= A[i] {
			x := i - stack[len(stack)-1]
			if x > res {
				res = x
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return res
}

func main() {
	A := []int{6, 0, 8, 2, 1, 5}
	fmt.Println(maxWidthRamp(A))
}
