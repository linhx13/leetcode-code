package main

import "fmt"

func isBipartite(graph [][]int) bool {
	visited := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if !dfs(graph, visited, i) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, visited []int, u int) bool {
	if visited[u] == 0 {
		visited[u] = 1
	}
	for _, v := range graph[u] {
		if visited[v] == 0 {
			visited[v] = -visited[u]
			if !dfs(graph, visited, v) {
				return false
			}
		} else {
			if visited[u] == visited[v] {
				return false
			}
		}
	}
	return true
}

func main() {
	// graph := [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}
	graph := [][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}}
	fmt.Println(isBipartite(graph))
}
