package main

var dirs = [][]int{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1},
	[]int{-1, -1}, []int{-1, 1}, []int{1, 1}, []int{1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dfs(board, click[0], click[1], m, n, visited)
	return board
}

func dfs(board [][]byte, i, j int, m, n int, visited [][]bool) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if visited[i][j] {
		return
	}
	visited[i][j] = true
	if board[i][j] == 'M' {
		board[i][j] = 'X'
		return
	}
	if board[i][j] == 'E' {
		c := mine(board, i, j, m, n)
		if c == 0 {
			board[i][j] = 'B'
			for k := 0; k < len(dirs); k++ {
				x := i + dirs[k][0]
				y := j + dirs[k][1]
				dfs(board, x, y, m, n, visited)
			}
		} else {
			board[i][j] = byte(c) + '0'
		}
	}
}

func mine(board [][]byte, i, j int, m, n int) int {
	cnt := 0
	for k := 0; k < len(dirs); k++ {
		x := i + dirs[k][0]
		y := j + dirs[k][1]
		if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'M' {
			cnt++
		}
	}
	return cnt
}
