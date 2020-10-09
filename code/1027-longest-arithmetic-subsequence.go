package main

import "fmt"

func longestArithSeqLength(A []int) int {
	dp := make([]map[int]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make(map[int]int)
	}
	res := 0
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			vj := 0
			if v, ok := dp[j][diff]; ok {
				vj = v
			} else {
				vj = 1
			}
			dp[i][diff] = max(dp[i][diff], 1+vj)
			res = max(res, dp[i][diff])
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
	A := []int{3, 6, 9, 12}
	fmt.Println(longestArithSeqLength(A))
}
