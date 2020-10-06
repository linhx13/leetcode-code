package main

import "fmt"

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for n > 0 {
		x := n % 10
		n = n / 10
		p *= x
		s += x
	}
	return p - s
}

func main() {
	fmt.Println(subtractProductAndSum(1))
}
