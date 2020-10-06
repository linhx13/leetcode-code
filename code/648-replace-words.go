package main

import "strings"

func replaceWords(dict []string, sentence string) string {
	m := make(map[string]bool)
	for _, w := range dict {
		m[w] = true
	}
	res := []string{}
	for _, word := range strings.Fields(sentence) {
		for i := 1; i <= len(word); i++ {
			v := m[word[0:i]]
			if v {
				word = word[0:i]
				break
			}
		}
		res = append(res, word)
	}
	return strings.Join(res, " ")
}
