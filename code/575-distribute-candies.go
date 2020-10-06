package main

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, c := range candies {
		m[c] = m[c] + 1
	}
	if len(m) <= len(candies)/2 {
		return len(m)
	} else {
		return len(candies) / 2
	}
}
