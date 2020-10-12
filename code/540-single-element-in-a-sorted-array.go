package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if mid-1 >= 0 && nums[mid] == nums[mid-1] {
			if (mid-low+1)%2 == 1 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
			if (mid+1-low+1)%2 == 1 {
				high = mid
			} else {
				low = mid + 2
			}
		} else {
			return nums[mid]
		}
	}
	return -1
}

func main() {
	// nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	// nums := []int{3, 3, 7, 7, 10, 11, 11}
	nums := []int{1, 2, 2}
	fmt.Println(singleNonDuplicate(nums))
}
