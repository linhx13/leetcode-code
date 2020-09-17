package main

import "fmt"

func maxScoreSightseeingPair(A []int) int {
	res := ^int(^uint(0) >> 1)
	left := A[0] + 0
	for j := 1; j < len(A); j++ {
		x := A[j] - j + left
		if x > res {
			res = x
		}
		c := A[j] + j
		if c > left {
			left = c
		}
	}
	return res
}

func main() {
	A := []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(A))
}
