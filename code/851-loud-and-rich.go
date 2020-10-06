package main

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	res := make([]int, n)
	people := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = []int{}
		res[i] = -1
		people[quiet[i]] = i
	}
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}
	for i := 0; i < n; i++ {
		res[i] = dfs(i, g, richer, quiet, &res)
	}
	for i := 0; i < n; i++ {
		res[i] = people[res[i]]
	}
	return res
}

func dfs(idx int, g [][]int, richer [][]int, quiet []int, res *[]int) int {
	if (*res)[idx] >= 0 {
		return (*res)[idx]
	}
	minn := quiet[idx]
	for i := 0; i < len(g[idx]); i++ {
		minn = min(dfs(g[idx][i], g, richer, quiet, res), minn)
	}
	(*res)[idx] = minn
	return minn
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
