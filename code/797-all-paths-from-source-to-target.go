package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	if len(graph) == 0 {
		return res
	}
	visited := make([]bool, len(graph))
	dfs(graph, 0, visited, &[]int{}, &res)
	return res
}

func dfs(graph [][]int, u int, visited []bool, path *[]int, res *[][]int) {
	if u == len(graph)-1 {
		*path = append(*path, u)
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
		*path = (*path)[0 : len(*path)-1]
		return
	}
	visited[u] = true
	*path = append(*path, u)
	for _, v := range graph[u] {
		if !visited[v] {
			dfs(graph, v, visited, path, res)
		}
	}
	*path = (*path)[0 : len(*path)-1]
	visited[u] = false
}

func main() {
	graph := [][]int{[]int{1, 2}, []int{3}, []int{3}, []int{}}
	fmt.Println(allPathsSourceTarget(graph))
}
