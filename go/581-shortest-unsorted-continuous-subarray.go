package main

func findUnsortedSubarray(nums []int) int {
	res := 0
	start := -1
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			j := i
			for j > 0 && nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				j--
			}
			if start == -1 || start > j {
				start = j
			}
			res = i - start + 1
		}
	}
	return res
}
