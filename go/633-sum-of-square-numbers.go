package main

import (
	"fmt"
	"math"
)

func squareCheck(a int) bool {
	var intRoot int = int(math.Sqrt(float64(a)))
	return (intRoot * intRoot) == a
}

func judgeSquareSum(c int) bool {
	for i := 0; i <= int(math.Sqrt(float64(c)/2)); i++ {
		if squareCheck(c - i*i) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(2147483644))
}
