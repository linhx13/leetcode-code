package main

import "sort"

func arrayRankTransform(arr []int) []int {
	rank := make(map[int]int)
	for _, n := range arr {
		rank[n] = 0
	}
	tmp := []int{}
	for k, _ := range rank {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	for i, n := range tmp {
		rank[n] = i + 1
	}
	res := []int{}
	for _, n := range arr {
		res = append(res, rank[n])
	}
	return res
}
