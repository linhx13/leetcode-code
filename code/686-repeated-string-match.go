package main

import "strings"

func repeatedStringMatch(A string, B string) int {
	s := A
	n := len(B)/len(A) + 2
	for i := 1; i <= n; i++ {
		if strings.Index(s, B) != -1 {
			return i
		}
		s += A
	}
	return -1
}
