package main

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ones_i := bits.OnesCount(uint(arr[i]))
		ones_j := bits.OnesCount(uint(arr[j]))
		if ones_i == ones_j {
			return arr[i] < arr[j]
		} else {
			return ones_i < ones_j
		}
	})
	return arr
}
