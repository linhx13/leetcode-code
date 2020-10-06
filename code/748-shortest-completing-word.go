package main

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	l := make([]int, 26)
	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			l[unicode.ToLower(c)-'a']++
		}
	}
	var res string
	min_l := int(^uint(0) >> 1)
	for _, word := range words {
		if len(word) >= min_l {
			continue
		}
		if !matches(l, word) {
			continue
		}
		min_l = len(word)
		res = word
	}
	return res
}

func matches(l []int, word string) bool {
	c := make([]int, 26)
	for _, r := range word {
		c[unicode.ToLower(r)-'a']++
	}
	for i := 0; i < 26; i++ {
		if c[i] < l[i] {
			return false
		}
	}
	return true
}
