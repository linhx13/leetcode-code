package main

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	edges_r := make([]map[int]bool, n)
	edges_b := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		edges_r[i] = make(map[int]bool)
		edges_b[i] = make(map[int]bool)
	}
	for _, e := range red_edges {
		edges_r[e[0]][e[1]] = true
	}
	for _, e := range blue_edges {
		edges_b[e[0]][e[1]] = true
	}
	seen_r := make(map[int]bool)
	seen_b := make(map[int]bool)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	queue := [][]int{}
	queue = append(queue, []int{0, 0})
	queue = append(queue, []int{0, 1})
	seen_b[0] = true
	seen_r[0] = true
	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			a := queue[0]
			queue = queue[1:len(queue)]
			p, is_red := a[0], a[1]
			if res[p] >= 0 {
				res[p] = min(res[p], steps)
			} else {
				res[p] = steps
			}
			edges := edges_b
			seen := seen_b
			if is_red != 0 {
				edges = edges_r
				seen = seen_r
			}
			for next, _ := range edges[p] {
				if seen[next] {
					continue
				}
				seen[next] = true
				queue = append(queue, []int{next, 1 - is_red})
			}
		}
		steps++
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
