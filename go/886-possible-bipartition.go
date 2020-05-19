package main

func dfs(g [][]int, cur int, color int, colors *[]int) bool {
	(*colors)[cur] = color
	for i := 0; i < len(g); i++ {
		if g[cur][i] == 1 {
			if (*colors)[i] == color {
				return false
			}
			if (*colors)[i] == 0 && !dfs(g, i, -color, colors) {
				return false
			}
		}
	}
	return true
}

func possibleBipartition(N int, dislikes [][]int) bool {
	g := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		g[i] = make([]int, N+1)
	}
	for _, d := range dislikes {
		g[d[0]][d[1]] = 1
		g[d[1]][d[0]] = 1
	}
	colors := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if colors[i] == 0 && !dfs(g, i, 1, &colors) {
			return false
		}
	}
	return true
}
