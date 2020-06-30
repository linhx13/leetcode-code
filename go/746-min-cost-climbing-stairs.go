package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i:=2;i<n+1;i++ {
		x := dp[i-2] + cost[i-2]
		y := dp[i-1] + cost[i-1]
		if x < y {
			dp[i] = x
		} else {
			dp[i] = y
		}
	}
	return dp[n]
}
