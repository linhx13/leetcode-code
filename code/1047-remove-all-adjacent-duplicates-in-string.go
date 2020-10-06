package main

import "fmt"

func removeDuplicates(S string) string {
	buf := []rune{}
	for _, c := range S {
		buf = append(buf, c)
		for {
			n := len(buf)
			if n >= 2 && buf[n-1] == buf[n-2] {
				buf = buf[:n-2]
			} else {
				break
			}
		}
	}
	return string(buf)
}

func main() {
	// S := "abbaca"
	S := "aaaaa"
	fmt.Println(removeDuplicates(S))
}
