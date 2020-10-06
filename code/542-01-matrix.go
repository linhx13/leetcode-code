package main

type Pair struct {
	X int
	Y int
}

func updateMatrix(matrix [][]int) [][]int {
	const INT_MAX = int(^uint(0) >> 1)
	r, c := len(matrix), len(matrix[0])
	dirs := []Pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	queue := []Pair{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, Pair{i, j})
			} else {
				matrix[i][j] = INT_MAX
			}
		}
	}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			x, y := p.X+d.X, p.Y+d.Y
			if x < 0 || x >= r || y < 0 || y >= c || matrix[x][y] <= matrix[p.X][p.Y] {
				continue
			}
			matrix[x][y] = matrix[p.X][p.Y] + 1
			queue = append(queue, Pair{x, y})
		}
	}
	return matrix
}
