package main

func maxDistance(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	res := -1
	q := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				q = append(q, 100*i+j)
			}
		}
	}
	dirs := [][]int{[]int{0, -1}, []int{0, 1}, []int{-1, 0}, []int{1, 0}}
	steps := 0
	for len(q) > 0 {
		for size := len(q); size > 0; size-- {
			p := q[0]
			q = q[1:len(q)]
			i, j := p/100, p%100
			if grid[i][j] == 2 {
				if steps > res {
					res = steps
				}
			}
			for k := 0; k < len(dirs); k++ {
				ii, jj := i+dirs[k][0], j+dirs[k][1]
				if ii < 0 || ii >= n || jj < 0 || jj >= m || grid[ii][jj] != 0 {
					continue
				}
				grid[ii][jj] = 2
				q = append(q, ii*100+jj)
			}
		}
		steps++
	}
	return res
}
