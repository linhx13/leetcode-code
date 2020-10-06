package main

import "fmt"

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}
	for j := 1; j <= N; j++ {
		dp[1][j] = j
	}
	for i := 2; i <= K; i++ {
		s := 1
		for j := 1; j <= N; j++ {
			dp[i][j] = j
			for s < j && dp[i-1][s-1] < dp[i][j-s] {
				s++
			}
			dp[i][j] = min(dp[i][j], max(dp[i-1][s-1], dp[i][j-s])+1)
		}
	}
	return dp[K][N]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	K, N := 3, 14
	fmt.Println(superEggDrop(K, N))
}
