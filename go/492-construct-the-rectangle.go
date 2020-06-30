package main

import (
	"math"
)

func constructRectangle(area int) []int {
	x := int(math.Ceil(math.Sqrt(float64(area))))
	for area%x != 0 {
		x++
	}
	y := area / x
	return []int{x, y}
}
