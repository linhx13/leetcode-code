package main

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	dp := make([]int, 100001)
	for i := 0; i < len(profit); i++ {
		dp[difficulty[i]] = max(profit[i], dp[difficulty[i]])
	}
	for i := 1; i <= 100000; i++ {
		dp[i] = max(dp[i], dp[i-1])
	}
	res := 0
	for i := 0; i < len(worker); i++ {
		res += dp[worker[i]]
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
