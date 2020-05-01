package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var r float64 = 1
	pos := true
	if n < 0 {
		pos = false
		n = -n
	}
	for i := n; i != 0; i >>= 1 {
		if i&1 == 1 {
			r *= x
		}
		x *= x
	}
	if !pos {
		r = 1.0 / r
	}
	return r
}

func main() {
	fmt.Println(myPow(-2, -2))
}
