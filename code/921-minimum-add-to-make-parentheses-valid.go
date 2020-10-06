package main

import "fmt"

func minAddToMakeValid(S string) int {
	stack := []rune{}
	for _, c := range S {
		if c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}
	return len(stack)
}

func main() {
	// s := "())"
	// s := "((("
	// s := "()"
	// s := "()))(("
	s := ")("
	fmt.Println(minAddToMakeValid(s))
}
