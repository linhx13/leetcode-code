package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		n = sum
		if m[n] {
			break
		}
		m[n] = true
	}
	return n == 1
}

func main() {
	n := 4
	fmt.Println(isHappy(n))
}
