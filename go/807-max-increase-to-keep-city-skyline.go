package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rows := make([]int, n)
	cols := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > rows[i] {
				rows[i] = grid[i][j]
			}
			if grid[i][j] > cols[j] {
				cols[j] = grid[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := rows[i]
			if x > cols[j] {
				x = cols[j]
			}
			res += x - grid[i][j]
		}
	}
	return res
}
