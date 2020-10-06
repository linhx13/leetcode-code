package main

import "fmt"

func numMatchingSubseq(S string, words []string) int {
	m := make(map[byte][]int)
	for i := 0; i < len(S); i++ {
		m[S[i]] = append(m[S[i]], i)
	}
	res := 0
	for i := 0; i < len(words); i++ {
		if helper(m, words[i]) {
			res++
		}
	}
	return res
}

func helper(m map[byte][]int, word string) bool {
	prev := -1
	for i := 0; i < len(word); i++ {
		c := word[i]
		hit := false
		for j := 0; j < len(m[c]); j++ {
			if m[c][j] > prev {
				prev = m[c][j]
				hit = true
				break
			}
		}
		if !hit {
			return false
		}
	}
	return true
}

func main() {
	S := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	fmt.Println(numMatchingSubseq(S, words))
}
