package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	freqs := make(map[int]int)
	for _, n := range arr {
		freqs[n]++
	}
	cnts := [][]int{}
	for k, v := range freqs {
		cnts = append(cnts, []int{k, v})
	}
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i][1] < cnts[j][1]
	})
	i := 0
	for i < len(cnts) {
		if cnts[i][1] <= k {
			k -= cnts[i][1]
			i++
		} else {
			break
		}
	}
	return len(cnts) - i
}

func main() {
	// arr := []int{5, 5, 4}
	// k := 1
	arr := []int{4, 3, 1, 1, 3, 3, 2}
	k := 3
	fmt.Println(findLeastNumOfUniqueInts(arr, k))
}
