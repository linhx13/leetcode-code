package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	idx_map := make(map[int]int)
	for idx, x := range arr2 {
		idx_map[x] = idx
	}
	sort.Slice(arr1, func(i int, j int) bool {
		idx_i, ok_i := idx_map[arr1[i]]
		idx_j, ok_j := idx_map[arr1[j]]
		if ok_i && ok_j {
			return idx_i < idx_j
		} else if ok_i {
			return true
		} else if ok_j {
			return false
		} else {
			return arr1[i] < arr1[j]
		}
	})
	return arr1
}
