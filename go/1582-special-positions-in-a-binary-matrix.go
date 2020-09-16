package main

func numSpecial(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	rowSums := make([]int, n)
	colSums := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rowSums[i] += mat[i][j]
			colSums[j] += mat[i][j]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 && rowSums[i] == 1 && colSums[j] == 1 {
				res++
			}
		}
	}
	return res
}
