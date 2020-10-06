package main

func projectionArea(grid [][]int) int {
	n := len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		rowMax, colMax := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res++
			}
			if grid[i][j] > rowMax {
				rowMax = grid[i][j]
			}
			if grid[j][i] > colMax {
				colMax = grid[j][i]
			}
		}
		res += rowMax + colMax
	}
	return res
}
