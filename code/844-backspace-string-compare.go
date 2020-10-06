package main

import "fmt"

func backspace(s string) string {
	buf := []rune{}
	for _, c := range s {
		if c == '#' {
			if len(buf) > 0 {
				buf = buf[:len(buf)-1]
			}
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

func backspaceCompare(S string, T string) bool {
	S = backspace(S)
	T = backspace(T)
	return S == T
}

func main() {
	// S, T := "ab#c", "ad#c"
	// S, T := "ab##", "c#d#"
	// S, T := "a##c", "#a#c"
	S, T := "a#c", "b"
	fmt.Println(backspaceCompare(S, T))
}
