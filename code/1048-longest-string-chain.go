package main

import (
	"fmt"
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	m := make(map[string]bool)
	for _, w := range words {
		m[w] = true
	}
	dp := make(map[string]int)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		dp[word] = max(0, dp[word])
		for k := 0; k < len(word); k++ {
			key := word[:k] + word[k+1:len(word)]
			if _, ok := m[key]; !ok {
				continue
			}
			dp[key] = max(dp[key], dp[word]+1)
		}
	}
	res := 0
	for _, v := range dp {
		if v+1 > res {
			res = v + 1
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	// words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	words := []string{"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh", "zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq", "gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx", "gru"}
	fmt.Println(longestStrChain(words))
}
