package main

func checkPossibility(nums []int) bool {
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if cnt == 0 {
				return false
			}
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
			cnt--
		}
	}
	return true
}
