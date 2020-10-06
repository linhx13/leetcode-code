package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	const INT_MAX = int(^uint(0) >> 1)
	min_abs := INT_MAX
	for i := 1; i < len(arr); i++ {
		abs := int(math.Abs(float64(arr[i] - arr[i-1])))
		if abs < min_abs {
			min_abs = abs
		}
	}
	res := [][]int{}
	for i := 1; i < len(arr); i++ {
		abs := int(math.Abs(float64(arr[i] - arr[i-1])))
		if abs == min_abs {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
