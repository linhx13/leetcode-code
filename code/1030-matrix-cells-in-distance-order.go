package main

import (
	"math"
	"sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		di := int(math.Abs(float64(res[i][0]-r0)) + math.Abs(float64(res[i][1]-c0)))
		dj := int(math.Abs(float64(res[j][0]-r0)) + math.Abs(float64(res[j][1]-c0)))
		return di < dj
	})
	return res
}
