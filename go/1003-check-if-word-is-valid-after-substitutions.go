package main

import "fmt"

func isValid(S string) bool {
	buf := []rune{}
	for _, c := range S {
		buf = append(buf, c)
		for {
			n := len(buf)
			if n >= 3 && buf[n-3] == 'a' && buf[n-2] == 'b' && buf[n-1] == 'c' {
				buf = buf[:n-3]
			} else {
				break
			}
		}
	}
	return len(buf) == 0
}

func main() {
	// s := "aabcbc"
	// s := "abcabcababcc"
	// s := "abccba"
	s := "cababc"
	fmt.Println(isValid(s))
}
