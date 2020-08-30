package main

import "fmt"

func expressiveWords(S string, words []string) int {
	res := 0
	for _, w := range words {
		if isValid(S, w) {
			res++
		}
	}
	return res
}

func isValid(s string, w string) bool {
	i, j := 0, 0
	var lastS, lastW byte
	var cntS, cntW int
	for i < len(s) && j < len(w) {
		for ; i < len(s); i++ {
			if s[i] != lastS && lastS != 0 {
				break
			}
			lastS = s[i]
			cntS++
		}
		for ; j < len(w); j++ {
			if w[j] != lastW && lastW != 0 {
				break
			}
			lastW = w[j]
			cntW++
		}
		if lastS != lastW {
			return false
		}
		if cntS != cntW && (cntS < cntW || cntS < 3) {
			return false
		}
		lastS, lastW = 0, 0
		cntS, cntW = 0, 0
	}
	if i < len(s) || j < len(w) {
		return false
	}
	return true
}

func main() {
	S := "heeellooo"
	words := []string{"hello", "hi", "helo", ""}
	fmt.Println(expressiveWords(S, words))
}
