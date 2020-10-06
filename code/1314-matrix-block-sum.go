package main

func matrixBlockSum(mat [][]int, K int) [][]int {
	n, m := len(mat), len(mat[0])
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for y := 1; y <= n; y++ {
		for x := 1; x <= m; x++ {
			dp[y][x] = dp[y-1][x] + dp[y][x-1] + mat[y-1][x-1] - dp[y-1][x-1]
		}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = mat[i][j]
		}
	}
	for y := 1; y <= n; y++ {
		for x := 1; x <= m; x++ {
			x1 := max(1, x-K)
			x2 := min(m, x+K)
			y1 := max(1, y-K)
			y2 := min(n, y+K)
			res[y-1][x-1] = dp[y2][x2] - dp[y1-1][x2] - dp[y2][x1-1] + dp[y1-1][x1-1]
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

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
