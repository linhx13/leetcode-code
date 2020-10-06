package main

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, x_center: x_center, y_center: y_center}
}

func (this *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64()) * this.radius
	t := rand.Float64() * 2 * math.Pi
	x := r*math.Cos(t) + this.x_center
	y := r*math.Sin(t) + this.y_center
	return []float64{x, y}
}
