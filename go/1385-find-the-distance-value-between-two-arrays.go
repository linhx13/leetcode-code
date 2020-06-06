package main

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, x := range arr1 {
		flag := true
		for _, y := range arr2 {
			if int(math.Abs(float64(x-y))) <= d {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}
