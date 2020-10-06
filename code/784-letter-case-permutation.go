package main

import (
	"fmt"
	"unicode"
)

func helper(runes []rune, idx int, res *[]string) {
	if idx == len(runes) {
		*res = append(*res, string(runes))
		return
	}
	helper(runes, idx+1, res)
	if !unicode.IsLetter(runes[idx]) {
		return
	}
	runes[idx] ^= (1 << 5)
	helper(runes, idx+1, res)
	runes[idx] ^= (1 << 5)
}

func letterCasePermutation(S string) []string {
	res := []string{}
	helper([]rune(S), 0, &res)
	return res
}

func main() {
	S := "a1b2"
	fmt.Println(letterCasePermutation(S))
}
