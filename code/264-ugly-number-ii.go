package main

import "fmt"

func nthUglyNumber(n int) int {
	i2, i3, i5 := 0, 0, 0
	mul2, mul3, mul5 := 2, 3, 5
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(min(mul2, mul3), mul5)
		if dp[i] == mul2 {
			i2++
			mul2 = dp[i2] * 2
		}
		if dp[i] == mul3 {
			i3++
			mul3 = dp[i3] * 3
		}
		if dp[i] == mul5 {
			i5++
			mul5 = dp[i5] * 5
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}
