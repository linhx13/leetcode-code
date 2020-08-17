package main

func minFallingPathSum(A [][]int) int {
	if len(A) == 1 {
		return A[0][0]
	}
	n := len(A)
	res := int(^uint(0) >> 1)
	for i := 1; i< n; i++ {
		for j:=0;j<n; j++{
			pre := A[i-1][j]
			if j > 0 {
				pre = min(pre, A[i-1][j-1])
			}
			if j < n-1 {
				pre = min(pre, A[i-1][j+1])
			}
			A[i][j] += pre
			if i == n -1 {
				res = min(res, A[i][j])
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
