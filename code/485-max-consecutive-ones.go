package main

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	last := 0
	cur := 0
	for _, n := range nums {
		if n == 0 {
			if last == 1 {
				if cur > res {
					res = cur
				}
			}
			cur, last = 0, 0
		} else {
			last = 1
			cur++
		}
	}
	if cur > res {
		res = cur
	}
	return res
}
