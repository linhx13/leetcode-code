package main

func stoneGameII(piles []int) int {
	n := len(piles)
	cache := make(map[int]int)
	total := 0
	for _, p := range piles {
		total += p
	}
	return (total + solve(0, 1, n, cache, piles)) / 2
}

func solve(s, M int, n int, cache map[int]int, piles []int) int {
	if s >= n {
		return 0
	}
	key := (s << 8) | M
	if v, ok := cache[key]; ok {
		return v
	}
	best := ^int(^uint(0) >> 1)
	cur := 0
	for x := 1; x <= 2*M; x++ {
		if s+x > n {
			break
		}
		cur += piles[s+x-1]
		best = max(best, cur-solve(s+x, max(x, M), n, cache, piles))
	}
	cache[key] = best
	return cache[key]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
