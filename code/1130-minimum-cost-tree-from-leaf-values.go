package main

import "fmt"

func mctFromLeafValues(arr []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	stack := []int{}
	stack = append(stack, INT_MAX)
	res := 0
	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1] < arr[i] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			s := stack[len(stack)-1]
			if arr[i] < s {
				s = arr[i]
			}
			res += t * s
		}
		stack = append(stack, arr[i])
	}
	for len(stack) > 2 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res += t * stack[len(stack)-1]
	}
	return res
}

func main() {
	nums := []int{6, 2, 4}
	fmt.Println(mctFromLeafValues(nums))
}
