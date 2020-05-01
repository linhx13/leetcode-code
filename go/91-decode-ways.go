package main

import "fmt"

func numDecodings(s string) int {
	runes := []rune(s)
	dp := make([]int, len(runes))
	for i := 0; i < len(runes); i++ {
		if runes[i] == '0' {
			if i == 0 {
				return 0
			} else if runes[i-1] == '0' || (runes[i-1] != '1' && runes[i-1] != '2') {
				return 0
			} else {
				if i-2 >= 0 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}
			}
		} else if (i-1 >= 0 && runes[i-1] == '1' && '1' <= runes[i] && runes[i] <= '9') || (i-1 >= 0 && runes[i-1] == '2' && '1' <= runes[i] && runes[i] <= '6') {
			if i-2 >= 0 {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = 2
			}
		} else {
			if i == 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(runes)-1]
}

func main() {
	fmt.Println(numDecodings("01"))
}
