package main

import "fmt"

func numRollsToTarget(d int, f int, target int) int {
	mod := int(1e9 + 7)
	dp := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int, target+1)
	}
	for j := 1; j <= f && j <= target; j++ {
		dp[1][j] = 1
	}
	for i := 1; i <= d; i++ {
		for j := 1; j <= target; j++ {
			for k := 1; k <= f; k++ {
				if i-1 >= 0 && j-k >= 0 && dp[i-1][j-k] > 0 {
					dp[i][j] += dp[i-1][j-k]
					if dp[i][j] > mod {
						dp[i][j] %= mod
					}
				}
			}
		}
	}
	return dp[d][target]
}

func main() {
	// d, f, target := 1, 6, 3
	d, f, target := 30, 30, 500
	fmt.Println(numRollsToTarget(d, f, target))
}
