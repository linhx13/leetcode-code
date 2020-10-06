package main

import (
	"fmt"
)

func countPrimes(n int) int {
	cnt := 0
	prime := make([]bool, n)
	for i := 0; i < n; i++ {
		prime[i] = true
	}
	for i := 2; i < n; i++ {
		if !prime[i] {
			continue
		}
		cnt++
		for j := 2; i*j < n; j++ {
			prime[i*j] = false
		}
	}
	return cnt
}

func main() {
	fmt.Println(countPrimes(10))
}
