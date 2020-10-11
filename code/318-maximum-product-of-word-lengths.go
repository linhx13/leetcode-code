package main

import "fmt"

func maxProduct(words []string) int {
	bm := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		x := 0
		for j := 0; j < len(words[i]); j++ {
			x |= 1 << int(words[i][j]-'a')
		}
		bm[i] = x
	}
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bm[i]&bm[j] == 0 {
				p := len(words[i]) * len(words[j])
				if p > res {
					res = p
				}
			}
		}
	}
	return res
}

func main() {
	// words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	// words := []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	words := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(maxProduct(words))
}
