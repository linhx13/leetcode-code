package main

func dfs(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '2' {
		return
	}
	grid[i][j] = '2'
	if i-1 >= 0 {
		dfs(grid, i-1, j)
	}
	if i+1 < len(grid) {
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 {
		dfs(grid, i, j-1)
	}
	if j+1 < len(grid[0]) {
		dfs(grid, i, j+1)
	}
}

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}
