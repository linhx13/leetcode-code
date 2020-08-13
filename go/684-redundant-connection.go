package main

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{parent: parent, rank: rank}
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DisjointSet) Union(x, y int) bool {
	px := ds.Find(x)
	py := ds.Find(y)
	if px == py {
		return false
	}

	if ds.rank[px] > ds.rank[py] {
		ds.parent[py] = px
	} else if ds.rank[px] < ds.rank[py] {
		ds.parent[px] = py
	} else {
		ds.parent[px] = py
		ds.rank[py]++
	}
	return true
}

func findRedundantConnection(edges [][]int) []int {
	ds := NewDisjointSet(len(edges) + 1)
	var res []int
	for _, edge := range edges {
		if !ds.Union(edge[0], edge[1]) {
			res = edge
		}
	}
	return res
}
