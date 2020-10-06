package main

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	r := int(math.Sqrt(float64(num)))
	for i := 2; i <= r; i++ {
		if num%i == 0 {
			x := num / i
			if i == x {
				x = 0
			}
			sum += i + x
		}
	}
	return sum == num
}
