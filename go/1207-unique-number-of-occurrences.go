package main

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for _, n := range arr {
		counts[n]++
	}
	cc := make(map[int]bool)
	for _, c := range counts {
		if cc[c] {
			return false
		}
		cc[c] = true
	}
	return true
}
