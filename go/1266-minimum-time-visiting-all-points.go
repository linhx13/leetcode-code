package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		diffX := int(math.Abs(float64(points[i-1][0]) - float64(points[i][0])))
		diffY := int(math.Abs(float64(points[i-1][1]) - float64(points[i][1])))
		if diffX < diffY {
			res += diffY
		} else {
			res += diffX
		}
	}
	return res
}
