package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	helper(nums, 0, &[]int{}, &res)
	return res
}

func helper(nums []int, i int, cur *[]int, res *[][]int) {
	if i >= len(nums) {
		t := make([]int, len(*cur))
		copy(t, *cur)
		*res = append(*res, t)
		return
	}
	helper(nums, i+1, cur, res)
	*cur = append(*cur, nums[i])
	helper(nums, i+1, cur, res)
	*cur = (*cur)[0 : len(*cur)-1]
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
}
