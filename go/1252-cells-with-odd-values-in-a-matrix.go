package main

func oddCells(n int, m int, indices [][]int) int {
	rows := make([]int, n)
	cols := make([]int, m)
	for _, idx := range indices {
		rows[idx[0]] ^= 1
		cols[idx[1]] ^= 1
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res += rows[i] ^ cols[j]
		}
	}
	return res
}
