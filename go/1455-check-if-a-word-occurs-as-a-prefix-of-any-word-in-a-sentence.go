package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	res := -1
	for idx, word := range strings.Fields(sentence) {
		if strings.HasPrefix(word, searchWord) {
			res = idx + 1
			break
		}
	}
	return res
}
