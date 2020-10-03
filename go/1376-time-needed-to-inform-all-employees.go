package main

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		if i != headID {
			children[manager[i]] = append(children[manager[i]], i)
		}
	}
	return dfs(children, headID, informTime)
}

func dfs(children [][]int, cur int, informTime []int) int {
	if len(children[cur]) == 0 {
		return 0
	}
	max := 0
	for _, c := range children[cur] {
		t := dfs(children, c, informTime)
		if t > max {
			max = t
		}
	}
	return max + informTime[cur]
}

func main() {
	n := 6
	headID := 2
	manager := []int{2, 2, -1, 2, 2, 2}
	informTime := []int{0, 0, 1, 0, 0, 0}
	fmt.Println(numOfMinutes(n, headID, manager, informTime))
}
