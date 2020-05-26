package main

func countNegatives(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	r, c := m-1, 0
	for r >= 0 && c < n {
		if grid[r][c] < 0 {
			res += n - c
			r--
		} else {
			c++
		}
	}
	return res
}
