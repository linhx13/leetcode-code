package main

func rotatedDigits(N int) int {
	res := 0
	for i := 1; i <= N; i++ {
		res += isValid(i)
	}
	return res
}

func isValid(n int) int {
	const INVALID_MASK = (1 << 3) | (1 << 4) | (1 << 7)
	const VALID_MASK = (1 << 2) | (1 << 5) | (1 << 6) | (1 << 9)

	valid := 0
	for n > 0 {
		r := 1 << (n % 10)
		if r&INVALID_MASK != 0 {
			return 0
		} else if r&VALID_MASK != 0 {
			valid = 1
		}
		n /= 10
	}
	return valid
}
