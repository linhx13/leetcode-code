package main

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	const INT_MAX = int(^uint(0) >> 1)
	a, b := INT_MAX, INT_MAX
	for i := 0; i < len(nums); i++ {
		if nums[i] <= a {
			a = nums[i]
		} else if nums[i] <= b {
			b = nums[i]
		} else {
			return true
		}
	}
	return false
}
