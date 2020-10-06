package main

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	min, max := -1, -1
	if arr[0] == arr[1]-1 && arr[1] == arr[2]-1 {
		min = 0
	} else if arr[1]-arr[0] <= 2 || arr[2]-arr[1] <= 2 {
		min = 1
	} else {
		min = 2
	}
	max = arr[2] - arr[1] - 1 + (arr[1] - arr[0] - 1)
	return []int{min, max}
}
