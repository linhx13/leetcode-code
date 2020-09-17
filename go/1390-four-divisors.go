package main

import (
	"fmt"
	"math"
)

func sumFourDivisors(nums []int) int {
	res := 0
	for _, n := range nums {
		res += helper(n)
	}
	return res
}

func helper(n int) int {
	s := 0
	c := 0
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			c++
			j := n / i
			if i == j {
				s += i
			} else {
				s += i + j
				c++
			}
			if c > 4 {
				break
			}
		}
	}
	if c == 4 {
		return s
	} else {
		return 0
	}
}

func main() {
	nums := []int{21, 4, 7}
	fmt.Println(sumFourDivisors(nums))
}
