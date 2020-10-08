package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	edges := make(map[int][]int)
	conns := make(map[string]bool)
	for _, c := range connections {
		edges[c[0]] = append(edges[c[0]], c[1])
		edges[c[1]] = append(edges[c[1]], c[0])
		k := fmt.Sprintf("%d-%d", c[0], c[1])
		conns[k] = true
	}

	res := 0
	visited := make([]bool, n)
	dfs(edges, conns, 0, visited, &res)
	return res
}

func dfs(edges map[int][]int, conns map[string]bool, cur int, visited []bool, res *int) {
	if visited[cur] {
		return
	}
	visited[cur] = true
	for _, i := range edges[cur] {
		if visited[i] {
			continue
		}
		k := fmt.Sprintf("%d-%d", i, cur)
		if !conns[k] {
			(*res)++
		}
		dfs(edges, conns, i, visited, res)
	}
}
