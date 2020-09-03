package main

import (
	"fmt"
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]int)
	lower := make(map[string]int)
	vowel := make(map[string]int)
	for i, w := range wordlist {
		exact[w] = i
		l := strings.ToLower(w)
		if _, ok := lower[l]; !ok {
			lower[l] = i
		}
		v := toVowel(w)
		if _, ok := vowel[v]; !ok {
			vowel[v] = i
		}
	}
	res := []string{}
	for _, q := range queries {
		l := strings.ToLower(q)
		v := toVowel(q)
		if _, ok := exact[q]; ok {
			res = append(res, q)
		} else if i, ok := lower[l]; ok {
			res = append(res, wordlist[i])
		} else if i, ok := vowel[v]; ok {
			res = append(res, wordlist[i])
		} else {
			res = append(res, "")
		}
	}
	return res
}

func toVowel(s string) string {
	res := []rune{}
	for _, ch := range strings.ToLower(s) {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			res = append(res, '#')
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func main() {
	// wordlist := []string{"KiTe", "kite", "hare", "Hare"}
	// queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	wordlist := []string{"YellOw"}
	queries := []string{"yollow"}
	fmt.Println(spellchecker(wordlist, queries))
}
