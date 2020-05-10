package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	if len(popped) > len(pushed) {
		return false
	}
	stack := []int{}
	i, j := 0, 0
	for i < len(pushed) && j < len(popped) {
		stack = append(stack, pushed[i])
		i++
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == len(popped)
}

func main() {
	// pushed := []int{1, 2, 3, 4, 5}
	// popped := []int{4, 5, 3, 2, 1}
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
