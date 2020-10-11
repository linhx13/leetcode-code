package main

import (
	"fmt"
	"math"
)

func closestDivisors(num int) []int {
	d1 := helper(num + 1)
	d2 := helper(num + 2)
	diff1 := int(math.Abs(float64(d1[0] - d1[1])))
	diff2 := int(math.Abs(float64(d2[0] - d2[1])))
	if diff1 < diff2 {
		return d1
	} else {
		return d2
	}
}

func helper(num int) []int {
	sqrt := int(math.Sqrt(float64(num)))
	for i := sqrt; i >= 1; i-- {
		if num%i == 0 {
			j := num / i
			return []int{i, j}
		}
	}
	return []int{1, num}
}

func main() {
	num := int(1e9)
	fmt.Println(num)
	fmt.Println(closestDivisors(num))
}
