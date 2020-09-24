package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	sort.Ints(coins)
	for _, c := range coins {
		if c < len(dp) {
			dp[c] = 1
		}
	}
	for i := 0; i <= amount; i++ {
		for _, c := range coins {
			if c > i {
				break
			}
			if dp[i-c] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[i-c] + 1
			} else {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	// coins := []int{1, 2, 5}
	// amount := 11
	// coins := []int{2}
	// amount := 1
	coins := []int{474, 83, 404, 3}
	amount := 264
	fmt.Println(coinChange(coins, amount))
}
