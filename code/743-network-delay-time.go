package main

func networkDelayTime(times [][]int, N int, K int) int {
	edges := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		edges[i] = make([]int, N+1)
		for j := 1; j <= N; j++ {
			edges[i][j] = -1
		}
	}
	for _, x := range times {
		edges[x[0]][x[1]] = x[2]
	}
	dist := make([]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = -1
	}
	dist[K] = 0
	queue := []int{K}
	for len(queue) > 0 {
		visited := make(map[int]bool)
		for i := len(queue); i > 0; i-- {
			u := queue[0]
			queue = queue[1:]
			for v := 1; v <= N; v++ {
				if edges[u][v] != -1 && (dist[v] == -1 || dist[u]+edges[u][v] < dist[v]) {
					if !visited[v] {
						visited[v] = true
						queue = append(queue, v)
					}
					dist[v] = dist[u] + edges[u][v]
				}
			}
		}
	}
	res := 0
	for i := 1; i <= N; i++ {
		if dist[i] == -1 {
			res = dist[i]
			break
		}
		if dist[i] > res {
			res = dist[i]
		}
	}
	return res
}
