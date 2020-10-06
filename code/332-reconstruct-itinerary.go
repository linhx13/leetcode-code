package main

import (
	"fmt"
	"sort"
)

func dfs(m map[string][]string, total int, cur string, res []string, visited map[string]bool) []string {
	res = append(res, cur)
	if len(res) == total {
		return res
	}
	for idx, v := range m[cur] {
		key := fmt.Sprintf("%s-%s-%d", cur, v, idx)
		if visited[key] {
			continue
		}
		visited[key] = true
		res = dfs(m, total, v, res, visited)
		if len(res) == total {
			return res
		}
		visited[key] = false
	}
	res = res[:len(res)-1]
	return res
}

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	for _, item := range tickets {
		m[item[0]] = append(m[item[0]], item[1])
	}
	for _, v := range m {
		sort.Strings(v)
	}
	res := []string{}
	visited := make(map[string]bool)
	return dfs(m, 1+len(tickets), "JFK", res, visited)
}
