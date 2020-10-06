package main

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {
	sz := len(cells)
	bm := 0
	mask := (1 << (sz - 1)) - 2
	m := make(map[int]int)
	for i := 0; i < sz; i++ {
		bm += (cells[i] << (sz - 1 - i))
	}
	for i := 0; i < N; {
		bm = ^((bm >> 1) ^ (bm << 1)) & mask
		if v, ok := m[bm]; ok {
			i = N - (N-i-1)%(i-v)
			continue
		}
		m[bm] = i
		i++
	}
	res := make([]int, sz)
	for i := 0; i < sz; i++ {
		if bm&(1<<(sz-1-i)) != 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

func main() {
	// cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	// N := 7
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
	N := 1000000000
	fmt.Println(prisonAfterNDays(cells, N))
}
