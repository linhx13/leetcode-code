package main

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	for _, n := range A {
		sumA += n
	}
	sumB := 0
	for _, n := range B {
		sumB += n
	}
	diff := (sumA - sumB) / 2
	st := make(map[int]bool)
	for _, n := range A {
		st[n] = true
	}
	for _, n := range B {
		if st[n+diff] {
			return []int{n + diff, n}
		}
	}
	return []int{}
}
