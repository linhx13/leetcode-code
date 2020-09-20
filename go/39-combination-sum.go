package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	helper(candidates, target, 0, 0, &[]int{}, &res)
	return res
}

func helper(candidates []int, target int, idx int, sum int, buf *[]int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		x := make([]int, len(*buf))
		copy(x, *buf)
		*res = append(*res, x)
		return
	}
	for i := idx; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		sum += candidates[i]
		*buf = append(*buf, candidates[i])
		helper(candidates, target, i, sum, buf, res)
		sum -= candidates[i]
		*buf = (*buf)[0 : len(*buf)-1]
	}
}

func main() {
	// candidates := []int{2, 3, 6, 7}
	// target := 7
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
