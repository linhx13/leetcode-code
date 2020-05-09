package main

import "fmt"

func find132pattern(nums []int) bool {
	const INT_MAX = int(^uint(0) >> 1)
	n := len(nums)
	mn := INT_MAX
	for j := 0; j < n; j++ {
		if nums[j] < mn {
			mn = nums[j]
		}
		if mn == nums[j] {
			continue
		}
		for k := n - 1; k > j; k-- {
			if mn < nums[k] && nums[j] > nums[k] {
				return true
			}
		}
	}
	return false
}

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{3, 1, 4, 2}
	nums := []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(nums))
}
