package main

import "fmt"

func brokenCalc(X int, Y int) int {
	res := 0
	for Y > X {
		if Y%2 != 0 {
			Y++
			res++
		}
		Y >>= 1
		res++
	}
	if X == Y {
		return res
	}
	res += X - Y
	return res
}

func main() {
	// X, Y := 3, 10
	// X, Y := 5, 8
	X, Y := 2, 3
	fmt.Println(brokenCalc(X, Y))
}
