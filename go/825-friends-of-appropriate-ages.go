package main

import (
	"fmt"
)

func numFriendRequests(ages []int) int {
	MAX_AGE := 120
	counts := make([]int, MAX_AGE+1)
	for _, a := range ages {
		counts[a]++
	}
	res := 0
	for A := 1; A <= MAX_AGE; A++ {
		for B := int(float64(0.5)*float64(A)) + 7 + 1; B <= A; B++ {
			cb := counts[B]
			if A == B {
				cb--
			}
			res += counts[A] * cb
		}
	}
	return res
}

func main() {
	// ages := []int{16, 17, 18}
	// ages := []int{16, 16}
	ages := []int{20, 30, 100, 110, 120}
	fmt.Println(numFriendRequests(ages))
}
