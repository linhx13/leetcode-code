package main

import "fmt"

func isPathCrossing(path string) bool {
	visited := make(map[string]bool)
	x, y := 0, 0
	key := fmt.Sprintf("%d:%d", x, y)
	visited[key] = true
	for _, r := range path {
		switch r {
		case 'N':
			y++
		case 'E':
			x++
		case 'S':
			y--
		case 'W':
			x--
		}
		key := fmt.Sprintf("%d:%d", x, y)
		if visited[key] {
			return true
		}
		visited[key] = true
	}
	return false
}
