package main

import "fmt"

func isRobotBounded(instructions string) bool {
	// 0: north, 1: west, 2: south, 3: east
	dd := [][]int{[]int{0, 1}, []int{-1, 0}, []int{0, -1}, []int{1, 0}}
	dir := 0
	x, y := 0, 0
	for i := 0; i < len(instructions); i++ {
		switch instructions[i] {
		case 'G':
			x += dd[dir][0]
			y += dd[dir][1]
		case 'L':
			dir = (dir + 1) % 4
		case 'R':
			dir = (dir + 4 - 1) % 4
		}
	}
	return dir != 0 || (x == 0 && y == 0)
}

func main() {
	instructions := "GG"
	fmt.Println(isRobotBounded(instructions))
}
