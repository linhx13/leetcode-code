package main

import "sort"

func twoCitySchedCost(costs [][]int) int {
	res := 0
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			res += costs[i][0]
		} else {
			res += costs[i][1]
		}
	}
	return res
}
