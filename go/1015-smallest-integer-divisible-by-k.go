package main

import "fmt"

func smallestRepunitDivByK(K int) int {
	r := 0
	for l := 1; l <= K; l++ {
		r = (r*10 + 1) % K
		if r == 0 {
			return l
		}
	}
	return -1
}

func main() {
	K := 2
	fmt.Println(smallestRepunitDivByK(K))
}
