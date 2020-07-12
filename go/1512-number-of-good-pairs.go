package main

func numIdenticalPairs(nums []int) int {
	cnts := make(map[int]int)
	res := 0
	for _, n := range nums {
		res += cnts[n]
		cnts[n]++
	}
	return res
}
