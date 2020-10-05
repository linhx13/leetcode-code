package main

import (
	"fmt"
	"math"
)

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = i
		sqrt := int(float64(math.Sqrt(float64(i))))
		for j := sqrt; j >= 1; j-- {
			if i%j == 0 {
				c := j + dp[i/j]
				if c < dp[i] {
					dp[i] = c
				}
			}
		}
	}
	return dp[n]
}

func main() {
	n := 8
	fmt.Println(minSteps(n))
}
