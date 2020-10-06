package main

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
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

func findCircleNum(M [][]int) int {
	n := len(M)
	res := n
	dsu := NewDSU(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				if dsu.Find(i) != dsu.Find(j) {
					dsu.Union(i, j)
					res--
				}
			}
		}
	}
	return res
}
