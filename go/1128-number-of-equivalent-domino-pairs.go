package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := make(map[string]int)
	res := 0
	for _, d := range dominoes {
		key1 := fmt.Sprintf("%d-%d", d[0], d[1])
		key2 := fmt.Sprintf("%d-%d", d[1], d[0])
		res += cnt[key1]
		if key1 != key2 {
			res += cnt[key2]
		}
		cnt[key1]++
	}
	return res
}
