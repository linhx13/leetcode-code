package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	parts := strings.Split(s, " ")
	new_parts := make([]string, 0, len(parts))
	for _, p := range parts {
		r := []rune(p)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		new_parts = append(new_parts, string(r))
	}
	return strings.Join(new_parts, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
