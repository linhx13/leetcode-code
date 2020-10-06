package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			if nums[low] <= target || nums[high-1] > nums[mid] {
				high = mid
			} else {
				low = mid + 1
			}
		case nums[mid] < target:
			if nums[high-1] >= target || nums[low] < nums[mid] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	// nums := []int{7, 8, 1, 2, 3, 4, 5, 6}
	fmt.Println(search(nums, 0))
}
