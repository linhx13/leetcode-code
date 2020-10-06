package main

func pivotIndex(nums []int) int {
	if len(nums) <= 3 {
		return -1
	}
	preSums := make([]int, len(nums))
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		preSums[i] = preSums[i-1] + nums[i-1]
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if preSums[i] == sum-nums[i]-preSums[i] {
			return i
		}
	}
	return -1
}
