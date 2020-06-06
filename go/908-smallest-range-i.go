package main

import "math"

func smallestRangeI(A []int, K int) int {
	min, max := int(^uint(0)>>1), 0
	for _, n := range A {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	if int(math.Abs(float64(max-min))) <= 2*K {
		return 0
	}
	return (max - K) - (min + K)
}
