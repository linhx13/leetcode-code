package main

func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 'X' {
				continue
			}
			h := j-1 >= 0 && board[i][j-1] == 'X'
			v := i-1 >= 0 && board[i-1][j] == 'X'
			if !h && !v {
				res++
			}
		}
	}
	return res
}
