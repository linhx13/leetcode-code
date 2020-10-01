package main

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if grid[i][j] == 0 && visited[i][j] == false {
					dfs(grid, m, n, i, j, visited)
				}
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && visited[i][j] == false {
				dfs(grid, m, n, i, j, visited)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, m, n int, i int, j int, visited [][]bool) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 || visited[i][j] == true {
		return
	}
	visited[i][j] = true
	dfs(grid, m, n, i-1, j, visited)
	dfs(grid, m, n, i+1, j, visited)
	dfs(grid, m, n, i, j-1, visited)
	dfs(grid, m, n, i, j+1, visited)
}
