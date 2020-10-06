package main

func dfs(grid [][]int, visited [][]bool, i int, j int, area *int) {
	if visited[i][j] {
		return
	}
	visited[i][j] = true
	if grid[i][j] == 0 {
		return
	}
	*area++
	if j-1 >= 0 {
		dfs(grid, visited, i, j-1, area)
	}
	if j+1 < len(grid[i]) {
		dfs(grid, visited, i, j+1, area)
	}
	if i-1 >= 0 {
		dfs(grid, visited, i-1, j, area)
	}
	if i+1 < len(grid) {
		dfs(grid, visited, i+1, j, area)
	}
}

func maxAreaOfIsland(grid [][]int) int {
	visited := [][]bool{}
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}
	var max int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[i][j] {
				continue
			}
			var area int
			dfs(grid, visited, i, j, &area)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func main() {

}
