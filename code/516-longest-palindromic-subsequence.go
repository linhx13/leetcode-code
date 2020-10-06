package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	t := string(runes)
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if s[i-1] == t[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[n][n]
}

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}
