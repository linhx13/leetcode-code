package main

import "fmt"

func constructArray(n int, k int) []int {
	res := []int{}
	for i := 1; i < n-k+1; i++ {
		res = append(res, i)
	}
	sign := -1
	res = append(res, n)
	for i := k - 1; i >= 1; i-- {
		res = append(res, res[len(res)-1]+sign*i)
		sign *= -1
	}
	return res
}

func main() {
	n, k := 3, 2
	fmt.Println(constructArray(n, k))
}
