package main

func gardenNoAdj(N int, paths [][]int) []int {
	res := make([]int, N)
	graph := make([][]int, N)
	for _, path := range paths {
		graph[path[0]-1] = append(graph[path[0]-1], path[1]-1)
		graph[path[1]-1] = append(graph[path[1]-1], path[0]-1)
	}
	for i := 0; i < N; i++ {
		colors := make(map[int]bool)
		for _, nei := range graph[i] {
			colors[res[nei]] = true
		}
		for color := 1; color < 5; color++ {
			if !colors[color] {
				res[i] = color
				break
			}
		}
	}
	return res
}
