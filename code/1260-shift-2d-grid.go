package main

func shiftGrid(grid [][]int, k int) [][]int {
	nums := []int{}
	for _, row := range grid {
		nums = append(nums, row...)
	}
	k = k % len(nums)
	tmp := []int{}
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	for i := len(nums) - 1 - k; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
	res := [][]int{}
	for i := 0; i < len(grid); i++ {
		res = append(res, nums[i*len(grid[0]):(i+1)*len(grid[0])])
	}
	return res
}
