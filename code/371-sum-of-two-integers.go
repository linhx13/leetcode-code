package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
}
