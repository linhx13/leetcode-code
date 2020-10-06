package main

import "fmt"

func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	t, _ := dfs(0, adj, hasApple, visited)
	return t
}

func dfs(node int, adj [][]int, hasApple []bool, visited []bool) (int, bool) {
	if node < 0 || node >= len(adj) {
		return 0, false
	}
	visited[node] = true
	t := 0
	b := false
	for _, c := range adj[node] {
		if visited[c] {
			continue
		}
		ct, cb := dfs(c, adj, hasApple, visited)
		if cb {
			b = true
			t += ct + 2
		}
	}
	if hasApple[node] {
		b = true
	}
	return t, b
}

func main() {
	n := 7
	edges := [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 4}, []int{1, 5}, []int{2, 3}, []int{2, 6}}
	hasApple := []bool{false, false, true, false, true, true, false}
	fmt.Println(minTime(n, edges, hasApple))
}
