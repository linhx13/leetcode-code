package main

import "math"

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	best := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				a := dist(points[i], points[j])
				b := dist(points[i], points[k])
				c := dist(points[j], points[k])
				s := (a + b + c) / 2
				area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
				if area > best {
					best = area
				}
			}
		}
	}
	return best
}

func dist(p1 []int, p2 []int) float64 {
	p1_x, p1_y := float64(p1[0]), float64(p1[1])
	p2_x, p2_y := float64(p2[0]), float64(p2[1])
	return math.Sqrt((p1_x-p2_x)*(p1_x-p2_x) + (p1_y-p2_y)*(p1_y-p2_y))
}
