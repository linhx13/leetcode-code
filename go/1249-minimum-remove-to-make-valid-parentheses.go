package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	idxStack := []int{}
	parStack := []rune{}
	for i, c := range s {
		if c == '(' {
			idxStack = append(idxStack, i)
			parStack = append(parStack, c)
		} else if c == ')' {
			if len(parStack) > 0 && parStack[len(parStack)-1] == '(' {
				parStack = parStack[:len(parStack)-1]
				idxStack = idxStack[:len(idxStack)-1]
			} else {
				idxStack = append(idxStack, i)
				parStack = append(parStack, c)
			}
		}
	}
	res := []rune{}
	j := 0
	for i, c := range s {
		if j < len(idxStack) && i == idxStack[j] {
			j++
			continue
		}
		res = append(res, c)
	}
	return string(res)
}

func main() {
	// s := "lee(t(c)o)de)"
	// s := "a)b(c)d"
	// s := "))(("
	s := "(a(b(c)d)"
	fmt.Println(minRemoveToMakeValid(s))
}
