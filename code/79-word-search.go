package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, idx int, x int, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[idx] {
		return false
	}
	if idx == len(word)-1 {
		return true
	}
	flag := false
	tmp := board[x][y]
	board[x][y] = ' '
	flag = dfs(board, word, idx+1, x, y+1) || dfs(board, word, idx+1, x, y-1) || dfs(board, word, idx+1, x+1, y) || dfs(board, word, idx+1, x-1, y)
	board[x][y] = tmp
	return flag
}
