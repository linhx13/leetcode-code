package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if 1 > len(res) {
			res = s[i : i+1]
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			dp[i][i+1] = true
			if 2 > len(res) {
				res = s[i : i+2]
			}
		}
	}
	for l := 1; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l - 1
			if j >= len(s) {
				continue
			}
			if !dp[i][j] {
				continue
			}
			if i-1 >= 0 && j+1 < len(s) && s[i-1] == s[j+1] {
				dp[i-1][j+1] = true
				if l+2 > len(res) {
					res = s[i-1 : i-1+l+2]
				}
			}
		}
	}
	return res
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}
