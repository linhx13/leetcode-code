package main

import (
	"fmt"
	"math"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if int(math.Abs(float64(arr[i]-arr[j]))) > a {
					continue
				}
				if int(math.Abs(float64(arr[j]-arr[k]))) > b {
					continue
				}
				if int(math.Abs(float64(arr[i]-arr[k]))) > c {
					continue
				}
				res++
			}
		}
	}
	return res
}

func main() {
	// arr := []int{3, 0, 1, 1, 9, 7}
	// a, b, c := 7, 2, 3
	arr := []int{1, 1, 2, 2, 3}
	a, b, c := 0, 0, 1
	fmt.Println(countGoodTriplets(arr, a, b, c))
}
