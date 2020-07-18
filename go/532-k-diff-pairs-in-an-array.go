package main

func findPairs(nums []int, k int) int {
	counts := make(map[int]int)
	res := 0
	for _, n := range nums {
		counts[n] ++
	}
	for n, v := range counts {
		if  k == 0 && v > 1 {
			res++
		}
		if k > 0 && counts[n+k] > 0 {
			res++
		}
	}
	return res
}
