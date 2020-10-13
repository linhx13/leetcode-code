package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for i := 0; i < len(intervals); i++ {
		low := intervals[i][0]
		high := intervals[i][1]
		for i = i + 1; i < len(intervals); i++ {
			if high < intervals[i][0] {
				break
			}
			if high < intervals[i][1] {
				high = intervals[i][1]
			}
		}
		i--
		res = append(res, []int{low, high})
	}
	return res
}

func main() {
	// intervals := [][]int{{3, 7}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {4, 5}}
	// intervals := [][]int{{1, 4}, {0, 2}, {3, 5}}
	fmt.Println(merge(intervals))
}
