package main

func findErrorNums(nums []int) []int {
	total := len(nums) * (1 + len(nums)) / 2
	sum := 0
	idx := make(map[int]int)
	twice := 0
	for _, n := range nums {
		sum += n
		idx[n]++
		if idx[n] == 2 {
			twice = n
		}
	}
	missing := total - sum + twice
	return []int{twice, missing}
}
