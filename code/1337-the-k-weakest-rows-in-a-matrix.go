package main

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	pairs := [][]int{}
	for i := 0; i < len(mat); i++ {
		c := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 1 {
				break
			}
			c += mat[i][j]
		}
		pairs = append(pairs, []int{c, i})
	}
	sort.Slice(pairs, func(i int, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		} else {
			return pairs[i][0] < pairs[j][0]
		}
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, pairs[i][1])
	}
	return res
}
