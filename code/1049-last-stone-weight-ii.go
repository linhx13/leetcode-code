package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, s := range stones {
		sum += s
	}
	dp := make([]int, sum/2+1)
	for _, s := range stones {
		for p := sum / 2; p > 0; p-- {
			if p >= s {
				if dp[p-s]+s > dp[p] {
					dp[p] = dp[p-s] + s
				}
			}
		}
	}
	return sum - 2*dp[sum/2]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(stones))
}
