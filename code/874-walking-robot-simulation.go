package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	obs := make(map[string]bool)
	for _, o := range obstacles {
		key := fmt.Sprintf("%d#%d", o[0], o[1])
		obs[key] = true
	}
	dirX := []int{0, 1, 0, -1}
	dirY := []int{1, 0, -1, 0}
	res := 0
	x, y, idx := 0, 0, 0
	for _, cmd := range commands {
		if cmd == -1 {
			idx = (idx + 1) % 4
		} else if cmd == -2 {
			idx = (idx - 1 + 4) % 4
		} else {
			for ; cmd > 0; cmd-- {
				key := fmt.Sprintf("%d#%d", x+dirX[idx], y+dirY[idx])
				if obs[key] {
					break
				}
				x += dirX[idx]
				y += dirY[idx]
			}
		}
		d := x*x + y*y
		if d > res {
			res = d
		}
	}
	return res
}
