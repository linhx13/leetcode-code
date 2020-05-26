package main

func islandPerimeter(grid [][]int) int {
	res := 0
	r, c := len(grid), len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			if i > 0 && i-1 >= 0 && grid[i-1][j] == 1 {
				res -= 1
			}
			if i < r-1 && i+1 < r && grid[i+1][j] == 1 {
				res -= 1
			}
			if j > 0 && j-1 >= 0 && grid[i][j-1] == 1 {
				res -= 1
			}
			if j < c-1 && j+1 < c && grid[i][j+1] == 1 {
				res -= 1
			}
		}
	}
	return res
}
