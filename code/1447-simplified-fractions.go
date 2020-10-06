package main

import "fmt"

func simplifiedFractions(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(simplifiedFractions(4))
}
