package main

import "fmt"

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := []int{}
	for i, t := range T {
		for len(stack) > 0 {
			topIdx := stack[len(stack)-1]
			if T[topIdx] >= t {
				break
			}
			res[topIdx] = i - topIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for _, idx := range stack {
		res[idx] = 0
	}
	return res
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
