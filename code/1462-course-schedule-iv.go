package main

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	if len(queries) == 0 {
		return []bool{}
	}
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = []int{}
	}
	for _, e := range prerequisites {
		g[e[0]] = append(g[e[0]], e[1])
	}

	visited := make([]bool, n)
	cache := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		cache[i] = make(map[int]bool)
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(g, i, visited, &[]int{}, cache)
		}
	}

	res := []bool{}
	for _, q := range queries {
		res = append(res, cache[q[0]][q[1]])
	}
	return res
}

func dfs(g [][]int, u int, visited []bool, path *[]int, cache []map[int]bool) {
	visited[u] = true
	*path = append(*path, u)
	for _, v := range g[u] {
		for _, p := range *path {
			cache[p][v] = true
			for x, _ := range cache[v] {
				cache[p][x] = true
			}
		}
		if !visited[v] {
			dfs(g, v, visited, path, cache)
		}
	}
	*path = (*path)[0 : len(*path)-1]
}
