package main

func findDuplicate(nums []int) int {
    left, right := 1, len(nums)
	for left < right {
		mid := left + (right - left) >> 1
		cnt := 0
		for _, n := range nums {
			if n <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid+1
		} else {
			right = mid
		}
	}
	return right
}
