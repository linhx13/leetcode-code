package main

func minStartValue(nums []int) int {
	min := int(^uint(0) >> 1)
	sum := 0
	for _, n := range nums {
		sum += n
		if sum < min {
			min = sum
		}
	}
	if min > 0 {
		return 1
	} else {
		return 1 - min
	}
}
