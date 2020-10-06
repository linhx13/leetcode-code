package main

import "fmt"

func scoreOfParentheses(S string) int {
	scores := []int{}
	for _, c := range S {
		if c == '(' {
			scores = append(scores, -1)
		} else if c == ')' {
			sum, i := 0, len(scores)-1
			for ; i >= 0; i-- {
				if scores[i] == -1 {
					break
				} else {
					sum += scores[i]
				}
			}
			if i == len(scores)-1 {
				scores[i] = 1
			} else {
				scores[i] = 2 * sum
				scores = scores[:i+1]
			}
		}
	}
	res := 0
	for _, s := range scores {
		res += s
	}
	return res
}

func main() {
	// s := "()"
	// s := "(())"
	// s := "(()())"
	// s := "(()(()))"
	s := ""
	fmt.Println(scoreOfParentheses(s))
}
