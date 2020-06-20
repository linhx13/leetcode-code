package main

import (
	"fmt"
	"sort"
)

type pair struct {
	Num int
	Idx int
}

func findRelativeRanks(nums []int) []string {
	ranks := []pair{}
	for i, n := range nums {
		ranks = append(ranks, pair{Num: n, Idx: i})
	}
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Num > ranks[j].Num
	})
	res := make([]string, len(nums))
	for i := 0; i < len(ranks); i++ {
		if i == 0 {
			res[ranks[i].Idx] = "Gold Medal"
		} else if i == 1 {
			res[ranks[i].Idx] = "Silver Medal"
		} else if i == 2 {
			res[ranks[i].Idx] = "Bronze Medal"
		} else {
			res[ranks[i].Idx] = fmt.Sprintf("%d", i+1)
		}
	}
	return res
}
