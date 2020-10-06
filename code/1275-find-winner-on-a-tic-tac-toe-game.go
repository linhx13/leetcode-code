package main

func tictactoe(moves [][]int) string {
	grid := [][]int{}
	for i := 0; i < 3; i++ {
		grid = append(grid, make([]int, 3))
	}
	for i, m := range moves {
		if i%2 == 0 {
			grid[m[0]][m[1]] = 1
		} else {
			grid[m[0]][m[1]] = 2
		}
	}
	for i := 0; i < 3; i++ {
		if grid[i][0] == grid[i][1] && grid[i][1] == grid[i][2] {
			if grid[i][0] == 1 {
				return "A"
			} else if grid[i][0] == 2 {
				return "B"
			}
		}
		if grid[0][i] == grid[1][i] && grid[1][i] == grid[2][i] {
			if grid[0][i] == 1 {
				return "A"
			} else if grid[0][i] == 2 {
				return "B"
			}
		}
	}
	if (grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2]) || (grid[2][0] == grid[1][1] && grid[1][1] == grid[0][2]) {
		if grid[1][1] == 1 {
			return "A"
		} else if grid[1][1] == 2 {
			return "B"
		}
	}
	if len(moves) < 9 {
		return "Pending"
	} else {
		return "Draw"
	}
}
