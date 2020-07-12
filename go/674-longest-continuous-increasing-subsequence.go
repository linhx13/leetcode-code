package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, 1
	res := 0
	for start < len(nums) && end < len(nums) {
		if nums[end] > nums[end-1] {
			end++
		} else {
			if end-start > res {
				res = end - start
			}
			start, end = end, end+1
		}
	}
	if end-start > res {
		res = end - start
	}
	return res
}
