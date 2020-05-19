package main

func dfs(board [][]byte, r, c int, i, j int) {
	board[i][j] = '$'
	if i-1 >= 0 && board[i-1][j] == 'O' {
		dfs(board, r, c, i-1, j)
	}
	if i+1 < r && board[i+1][j] == 'O' {
		dfs(board, r, c, i+1, j)
	}
	if j-1 >= 0 && board[i][j-1] == 'O' {
		dfs(board, r, c, i, j-1)
	}
	if j+1 < c && board[i][j+1] == 'O' {
		dfs(board, r, c, i, j+1)
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	r, c := len(board), len(board[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 || i == r-1 || j == 0 || j == c-1 {
				if board[i][j] == 'O' {
					dfs(board, r, c, i, j)
				}
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '$' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {

}
