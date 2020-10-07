package main

import "fmt"

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	res := [][]int{}
	helper(nums, used, &[]int{}, &res)
	return res
}

func helper(nums []int, used []bool, buf *[]int, res *[][]int) {
	if len(*buf) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *buf)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		*buf = append(*buf, nums[i])
		used[i] = true
		helper(nums, used, buf, res)
		*buf = (*buf)[0 : len(*buf)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	res := permute(nums)
	fmt.Println(res)
	fmt.Println(len(res))
}
