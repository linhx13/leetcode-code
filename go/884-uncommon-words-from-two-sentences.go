package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	cntA := make(map[string]int)
	for _, word := range strings.Fields(A) {
		cntA[word] = cntA[word] + 1
	}
	cntB := make(map[string]int)
	for _, word := range strings.Fields(B) {
		cntB[word] = cntB[word] + 1
	}
	res := []string{}
	for word, cnt := range cntA {
		if cnt == 1 && cntB[word] == 0 {
			res = append(res, word)
		}
	}
	for word, cnt := range cntB {
		if cnt == 1 && cntA[word] == 0 {
			res = append(res, word)
		}
	}
	return res
}
