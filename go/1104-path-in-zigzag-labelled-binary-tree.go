package main

import "math"

func pathInZigZagTree(label int) []int {
	res := []int{}
	n := label
	for n != 1 {
		res = append(res, n)
		level := int(math.Floor(float64(math.Log2(float64(n)))))
		if level%2 == 0 {
			n = n / 2
			n = int(math.Pow(float64(2), float64(level-1))) + int(math.Pow(float64(2), float64(level))) - 1 - n
		}
		if level%2 == 1 {
			n = int(math.Pow(float64(2), float64(level))) + int(math.Pow(float64(2), float64(level+1))) - 1 - n
			n = n / 2
		}
	}
	res = append(res, 1)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
