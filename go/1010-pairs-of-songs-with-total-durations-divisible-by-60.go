package main

func numPairsDivisibleBy60(time []int) int {
	buckets := make([]int, 60)
	res := 0
	for _, t := range time {
		t %= 60
		res += buckets[(60-t)%60]
		buckets[t]++
	}
	return res
}
