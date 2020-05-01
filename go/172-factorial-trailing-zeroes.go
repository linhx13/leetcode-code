package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0
	f := 5
	for f <= n {
		res += n / f
		f *= 5
	}
	return res
}

func main() {
	fmt.Println(trailingZeroes(12345))
}
