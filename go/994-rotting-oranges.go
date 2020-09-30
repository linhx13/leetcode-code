package main

import "fmt"

type Item struct {
	i, j int
	t    int
}

func orangesRotting(grid [][]int) int {
	res := 0
	all := 0
	n, m := len(grid), len(grid[0])
	seen := make(map[int]bool)
	queue := []Item{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, Item{i: i, j: j, t: 0})
				seen[i<<8+j] = true
			}
			if grid[i][j] != 0 {
				all++
			}
		}
	}
	dirs := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:len(queue)]
		for _, d := range dirs {
			i, j := cur.i+d[0], cur.j+d[1]
			if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != 1 {
				continue
			}
			k := i<<8 + j
			if !seen[k] {
				queue = append(queue, Item{i: i, j: j, t: cur.t + 1})
				seen[k] = true
				if cur.t+1 > res {
					res = cur.t + 1
				}
			}
		}
	}
	if len(seen) != all {
		return -1
	} else {
		return res
	}
}

func main() {
	grid := [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
