package main

import "fmt"

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	dp := make([][]float64, K)
	for i := 0; i < K; i++ {
		dp[i] = make([]float64, n)
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += float64(A[i])
		dp[0][i] = sum / float64(i+1)
	}
	for i := 1; i < K; i++ {
		for j := i; j < n; j++ {
			dp[i][j] = dp[i-1][j-1] + float64(A[j])

			sum := float64(A[j])
			for k := j - 1; k >= i; k-- {
				sum += float64(A[k])
				x := dp[i-1][k-1] + sum/float64(j-k+1)
				if x > dp[i][j] {
					dp[i][j] = x
				}
			}
		}
	}
	return dp[K-1][n-1]
}

func main() {
	A := []int{9, 1, 2, 3, 9}
	K := 3
	fmt.Println(largestSumOfAverages(A, K))
}
