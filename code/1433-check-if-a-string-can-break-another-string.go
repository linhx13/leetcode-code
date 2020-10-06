package main

import (
	"fmt"
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	res := true
	for i := 0; i < len(runes1); i++ {
		if runes1[i] < runes2[i] {
			res = false
			break
		}
	}
	if res {
		return res
	}

	res = true
	for i := 0; i < len(runes1); i++ {
		if runes1[i] > runes2[i] {
			res = false
			break
		}
	}
	return res
}

func main() {
	// s1, s2 := "abc", "xya"
	// s1, s2 := "abe", "acd"
	s1, s2 := "leetcode", "interview"
	fmt.Println(checkIfCanBreak(s1, s2))
}
