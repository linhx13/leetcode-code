package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}
		m1 := make(map[byte]byte)
		m2 := make(map[byte]byte)
		valid := true
		for i := 0; i < len(pattern); i++ {
			a, b := word[i], pattern[i]
			va, ok_a := m1[a]
			vb, ok_b := m2[b]
			if ok_a != ok_b {
				valid = false
				break
			} else if ok_a && ok_b && !(va == b && vb == a) {
				valid = false
				break
			} else {
				m1[a] = b
				m2[b] = a
			}
		}
		if valid {
			res = append(res, word)
		}
	}
	return res
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
