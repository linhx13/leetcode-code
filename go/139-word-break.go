package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}
		for _, word := range wordDict {
			l := len(word)
			if i+l > n {
				continue
			}
			if word == s[i:i+l] {
				dp[i+l] = true
			}
		}
	}
	return dp[n]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
