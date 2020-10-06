package main

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	for i := 0; i < n; i ++ {
		parent[i] = i
	}
	return &DSU{parent: parent}
}

func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x])
	}
	return dsu.parent[x]
}

func (dsu *DSU) Union(x, y int) {
	dsu.parent[dsu.Find(x)] = dsu.Find(y)
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	dsu := NewDSU(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := 4 * (i * n + j)
			switch grid[i][j] {
			case '/':
				dsu.Union(idx+0, idx+3)
				dsu.Union(idx+1, idx+2)
			case '\\':
				dsu.Union(idx+0, idx+1)
				dsu.Union(idx+2, idx+3)
			case ' ':
				dsu.Union(idx+0, idx+1)
				dsu.Union(idx+1, idx+2)
				dsu.Union(idx+2, idx+3)
			}
			if i + 1 < n {
				dsu.Union(idx+2, idx+4*n+0)
			}
			if j + 1 < n {
				dsu.Union(idx+1, idx+4+3)
			}
		}
	}
	res := 0
	for i := 0; i<4*n*n; i++ {
		if dsu.Find(i) == i {
			res++
		}
	}
	return res
}
