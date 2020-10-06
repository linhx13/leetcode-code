package main

import "fmt"

func getFreq(w string) int {
	cnts := make([]int, 26)
	for _, r := range []rune(w) {
		cnts[r-'a']++
	}
	for _, c := range cnts {
		if c > 0 {
			return c
		}
	}
	return 0
}

func numSmallerByFrequency(queries []string, words []string) []int {
	ws := make([]int, len(words))
	for _, w := range words {
		ws = append(ws, getFreq(w))
	}
	res := []int{}
	for _, q := range queries {
		f := getFreq(q)
		c := 0
		for _, w := range ws {
			if f < w {
				c++
			}
		}
		res = append(res, c)
	}
	return res
}

func main() {
	// queries := []string{"cbd"}
	// words := []string{"zaaaz"}
	queries := []string{"bbb", "cc"}
	words := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(numSmallerByFrequency(queries, words))
}
