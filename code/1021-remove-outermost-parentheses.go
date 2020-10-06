package main

import "fmt"

func removeOuterParentheses(S string) string {
	prevIdx := 0
	res := []rune{}
	count := 0
	for idx, c := range S {
		if c == '(' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 {
			res = append(res, []rune(S[prevIdx+1:idx])...)
			prevIdx = idx + 1
		}
	}
	return string(res)
}

func main() {
	// S := "(()())(())"
	// S := "(()())(())(()(()))
	S := "()()"
	fmt.Println(removeOuterParentheses(S))
}
