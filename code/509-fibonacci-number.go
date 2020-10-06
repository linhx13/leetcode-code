package main

func fib(N int) int {
	if N == 0 {
		return 0
	}
	a, b := 0, 1
	for N > 1 {
		a, b = b, a+b
		N--
	}
	return b
}
