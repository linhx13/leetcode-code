package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	m := make(map[int]bool)
	for _, q := range queens {
		m[q[0]*10+q[1]] = true
	}
	res := [][]int{}
	dirs := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1},
		[]int{-1, -1}, []int{-1, 1}, []int{1, -1}, []int{1, 1}}
	for _, d := range dirs {
		x, y := king[0], king[1]
		for {
			x, y = x+d[0], y+d[1]
			if !(x >= 0 && x < 8 && y >= 0 && y < 8) {
				break
			}
			if m[x*10+y] {
				res = append(res, []int{x, y})
				break
			}
		}
	}
	return res
}
