package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	res := 0
	dp := make([]int, len(pairs))
	for i := 0; i < len(pairs); i++ {
		dp[i] = 1
		for j := 0; j < len(pairs); j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
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

func main() {
	pairs := [][]int{[]int{1, 2}, []int{2, 3}}
	// pairs := [][]int{[]int{-6, 9}, []int{1, 6}, []int{8, 10}, []int{-1, 4}, []int{-6, -2}, []int{-9, 8}, []int{-5, 3}, []int{0, 3}}
	// pairs := [][]int{[]int{-1, 1}, []int{-2, 7}, []int{-5, 8}, []int{-3, 8}, []int{1, 3}, []int{-2, 9}, []int{-5, 2}}
	fmt.Println(findLongestChain(pairs))
}
