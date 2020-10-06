package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = true
		adj[e[1]][e[0]] = true
	}
	queue := []int{}
	for i := 0; i < len(adj); i++ {
		if len(adj[i]) == 1 {
			queue = append(queue, i)
		}
	}
	for n > 2 {
		size := len(queue)
		n -= size
		for i := 0; i < size; i++ {
			t := queue[0]
			queue = queue[1:len(queue)]
			for a, _ := range adj[t] {
				delete(adj[a], t)
				if len(adj[a]) == 1 {
					queue = append(queue, a)
				}
			}
		}
	}
	res := []int{}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:len(queue)]
		res = append(res, t)
	}
	return res
}

func main() {

}
