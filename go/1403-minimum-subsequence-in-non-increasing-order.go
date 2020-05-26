package main

func minSubsequence(nums []int) []int {
	counts := make([]int, 101)
	half_sum := 0
	for _, n := range nums {
		counts[n]++
		half_sum += n
	}
	half_sum /= 2
	res := []int{}
	sum := 0
	for i := 100; i >= 1 && sum <= half_sum; i-- {
		for counts[i] > 0 && sum <= half_sum {
			sum += i
			res = append(res, i)
			counts[i]--
		}
	}
	return res
}
