package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	words := strings.Fields(s)
	maxLen := 0
	for _, w := range words {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}
	res := []string{}
	for j := 0; j < maxLen; j++ {
		cur := []byte{}
		for i := 0; i < len(words); i++ {
			if j < len(words[i]) {
				cur = append(cur, words[i][j])
			} else {
				cur = append(cur, ' ')
			}
		}
		res = append(res, strings.TrimRight(string(cur), " "))
	}
	return res
}

func main() {
	// s := "HOW ARE YOU"
	s := "TO BE OR NOT TO BE"
	// s := "CONTEST IS COMING"
	fmt.Println(printVertically(s))
}
