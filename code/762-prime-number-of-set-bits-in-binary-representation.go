package main

import "math/bits"

func countPrimeSetBits(L int, R int) int {
	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}
	res := 0
	for n := L; n <= R; n++ {
		ones := bits.OnesCount(uint(n))
		if primes[ones] {
			res++
		}
	}
	return res
}
