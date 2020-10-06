package main

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			if grid[i+1][j+1] == 5 && isValid(grid, i, j) {
				res++
			}
		}
	}
	return res
}

func isValid(grid [][]int, i, j int) bool {
	cnt := make([]int, 10)
	for x := i; x < i+2; x++ {
		for y := j; y < j+2; y++ {
			k := grid[x][y]
			if k < 1 || k > 9 || cnt[k] != 0 {
				return false
			}
			cnt[k] = 1
		}
	}
	if 15 != grid[i][j]+grid[i][j+1]+grid[i][j+2] {
		return false
	}
	if 15 != grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] {
		return false
	}
	if 15 != grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] {
		return false
	}
	if 15 != grid[i][j]+grid[i+1][j]+grid[i+2][j] {
		return false
	}
	if 15 != grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] {
		return false
	}
	if 15 != grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] {
		return false
	}
	if 15 != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] {
		return false
	}
	if 15 != grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] {
		return false
	}
	return true
}
