package main

func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i< n; i++ {
		n := start + 2 * i
		if i == 0 {
			res = n
		} else {
			res ^= n
		}
	}
	return res
}
