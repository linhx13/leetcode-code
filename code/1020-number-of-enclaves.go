package main

func numEnclaves(A [][]int) int {
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == m-1 || j == 0 {
				if A[i][j] == 1 {
					dfs(A, i, j)
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs(A [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(A) || j >= len(A[0]) || A[i][j] == 0 || A[i][j] == 2 {
		return
	}
	A[i][j] = 2
	dfs(A, i+1, j)
	dfs(A, i, j+1)
	dfs(A, i, j-1)
	dfs(A, i-1, j)
}
