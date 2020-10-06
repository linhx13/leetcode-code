package main

import "fmt"

func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)
	idx := -1
	for low < high {
		mid := low + (high-low)>>2
		if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			idx = mid
			break
		}
	}
	if idx == -1 {
		return []int{-1, -1}
	}
	min, max := idx, idx
	for i := idx; i >= 0 && nums[i] == nums[idx]; min, i = i, i-1 {
	}
	for i := idx; i < len(nums) && nums[i] == nums[idx]; max, i = i, i+1 {
	}
	return []int{min, max}
}

func main() {
	// nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1, 2, 2, 3, 3}
	fmt.Println(searchRange(nums, 3))
}
