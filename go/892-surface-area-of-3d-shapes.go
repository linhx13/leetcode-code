package main

func surfaceArea(grid [][]int) int {
	n := len(grid)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res += 4*grid[i][j] + 2
			}
			if i > 0 {
				if grid[i][j] < grid[i-1][j] {
					res -= grid[i][j] * 2
				} else {
					res -= grid[i-1][j] * 2
				}
			}
			if j > 0 {
				if grid[i][j] < grid[i][j-1] {
					res -= grid[i][j] * 2
				} else {
					res -= grid[i][j-1] * 2
				}
			}
		}
	}
	return res
}
