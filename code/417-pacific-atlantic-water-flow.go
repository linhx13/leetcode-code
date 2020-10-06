package main

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	n, m := len(matrix), len(matrix[0])
	p := make([][]int, n)
	a := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		p[i] = make([]int, m)
		a[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		dfs(matrix, 0, j, 0, p)
		dfs(matrix, n-1, j, 0, a)
	}
	for i := 0; i < n; i++ {
		dfs(matrix, i, 0, 0, p)
		dfs(matrix, i, m-1, 0, a)
	}
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if p[i][j] != 0 && a[i][j] != 0 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(matrix [][]int, i int, j int, k int, visited [][]int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return
	}
	if visited[i][j] != 0 || matrix[i][j] < k {
		return
	}
	visited[i][j] = 1
	dfs(matrix, i-1, j, matrix[i][j], visited)
	dfs(matrix, i+1, j, matrix[i][j], visited)
	dfs(matrix, i, j-1, matrix[i][j], visited)
	dfs(matrix, i, j+1, matrix[i][j], visited)
}
