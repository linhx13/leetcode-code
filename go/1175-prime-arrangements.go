package main

import "math"

func numPrimeArrangements(n int) int {
	const MOD = 1e9 + 7
	p := 0
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			p++
		}
	}
	var res int64 = 1
	for i := 1; i <= p; i++ {
		res = (res * int64(i)) % MOD
	}
	for i := 1; i <= n-p; i++ {
		res = (res * int64(i)) % MOD
	}
	return int(res)
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	sqrt := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrt; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
