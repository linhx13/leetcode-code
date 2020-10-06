package main

func getMoneyAmount(n int) int {
	dp := make([]int, n+1)
	for i:= 0; i<= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for l:=1; l <=n; l++ {
		for i := 0; i+l-1 <= n; i++ {
			j := i+l-1
			if l == 1 {
				dp[i][j] = 0
			} else if l == 2 {
				dp[i][j] = i
			} else {
				res := int(^uint(0)>> 1)
				for k := i+1; k < j; k++ {
					res = min(res, k + max(dp[i][k-1], dp[k+1][j]))
				}
				dp[i][j] = res
			}
		}
	}
	return dp[1][n]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
