package main

func numRookCaptures(board [][]byte) int {
	res := 0
	x, y := -1, -1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
				break
			}
		}
		if x != -1 && y != -1 {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] != '.' {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] != '.' {
			break
		}
	}
	for j := y - 1; j >= 0; j-- {
		if board[x][j] == 'p' {
			res++
			break
		} else if board[x][j] != '.' {
			break
		}
	}
	for j := y + 1; j < 8; j++ {
		if board[x][j] == 'p' {
			res++
			break
		} else if board[x][j] != '.' {
			break
		}
	}
	return res
}
