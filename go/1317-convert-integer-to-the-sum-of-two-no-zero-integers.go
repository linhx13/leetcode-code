package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		j := n - i
		if checkValid(i) && checkValid(j) {
			return []int{i, j}
		}
	}
	return nil
}

func checkValid(n int) bool {
	if n == 0 {
		return false
	}
	for n > 0 {
		c := n % 10
		if c == 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func main() {
	n := 10000
	fmt.Println(getNoZeroIntegers(n))
}
