package main

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		l := len(stones)
		sort.Ints(stones)
		s := stones[l-1] - stones[l-2]
		stones = stones[:l-2]
		if s > 0 {
			stones = append(stones, s)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}
