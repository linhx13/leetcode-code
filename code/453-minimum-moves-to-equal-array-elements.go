package main

func minMoves(nums []int) int {
	min := int(^uint(0) >> 1)
	sum := 0
	for _, n:= range nums {
		if n < min {
			min = n
		}
		sum += n
	}
	return sum - len(nums) * min
}
