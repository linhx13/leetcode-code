package main

func repeatedNTimes(A []int) int {
	cnt := make(map[int]int)
	for _, n := range A {
		cnt[n]++
		if cnt[n] > 1 {
			return n
		}
	}
	return 0
}
