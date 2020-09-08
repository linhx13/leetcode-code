package main

func diagonalSum(mat [][]int) int {
	res := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		res += mat[i][i]
		res += mat[i][n-i-1]
	}
	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}
