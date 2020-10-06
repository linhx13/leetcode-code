package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := int(^uint(0) >> 1)
	minDiff := int(^uint(0) >> 1)
	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := first + nums[start] + nums[end]
			diff := int(math.Abs(float64(target - sum)))
			if diff < minDiff {
				minDiff = diff
				res = sum
			}
			if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return res
}
