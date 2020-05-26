package main

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	row_min := make([]int, m)
	col_max := make([]int, n)
	for i := 0; i < m; i++ {
		row_min[i] = 999999
		for j := 0; j < n; j++ {
			if matrix[i][j] < row_min[i] {
				row_min[i] = matrix[i][j]
			}
			if matrix[i][j] > col_max[j] {
				col_max[j] = matrix[i][j]
			}
		}
	}
	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == row_min[i] && matrix[i][j] == col_max[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
