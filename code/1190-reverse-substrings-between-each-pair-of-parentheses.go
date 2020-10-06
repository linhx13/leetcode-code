package main

import "fmt"

func reverseParentheses(s string) string {
	buf := []rune{}
	stack := []int{}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, len(buf))
		} else if c == ')' {
			startIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if startIdx < len(buf) {
				n := len(buf) - startIdx
				for i := 0; i < n/2; i++ {
					buf[startIdx+i], buf[startIdx+n-i-1] = buf[startIdx+n-i-1], buf[startIdx+i]
				}
			}
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

func main() {
	// s := "(abcd)"
	// s := "(u(love)i)"
	// s := "(ed(et(oc))el)"
	// s := "a(bcdefghijkl(mno)p)q"
	s := "((a)()(b))"
	fmt.Println(reverseParentheses(s))
}
