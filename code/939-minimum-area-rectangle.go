package main

import "math"

func minAreaRect(points [][]int) int {
	m := make(map[int]map[int]bool)
	for _, p := range points {
		if _, ok := m[p[0]]; !ok {
			m[p[0]] = make(map[int]bool)
		}
		m[p[0]][p[1]] = true
	}
	INT_MAX := int(^uint(0) >> 1)
	n := len(points)
	min := INT_MAX
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x0, y0 := points[i][0], points[i][1]
			x1, y1 := points[j][0], points[j][1]
			if x0 == x1 || y0 == y1 {
				continue
			}
			area := int(math.Abs(float64((x0 - x1) * (y0 - y1))))
			if area > min {
				continue
			}
			if m[x1][y0] == false || m[x0][y1] == false {
				continue
			}
			min = area
		}
	}
	if min == INT_MAX {
		return 0
	} else {
		return min
	}
}
