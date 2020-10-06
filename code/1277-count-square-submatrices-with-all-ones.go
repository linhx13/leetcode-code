package main

func countSquares(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = matrix[i][j]
			if i != 0 && j != 0 && dp[i][j] != 0 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1])
				dp[i][j] = min(dp[i][j], dp[i-1][j])
				dp[i][j]++
			}
			res += dp[i][j]
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
