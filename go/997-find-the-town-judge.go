package main

func findJudge(N int, trust [][]int) int {
	degree := make([]int, 1+N)
	flag := make([]int, 1+N)
	for _, pair := range trust {
		flag[pair[0]] = -1
		degree[pair[1]] += 1
	}
	for i := 1; i < N+1; i++ {
		if degree[i] == N-1 && flag[i] != -1 {
			return i
		}
	}
	return -1
}
