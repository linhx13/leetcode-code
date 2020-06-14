package main

func hasAlternatingBits(n int) bool {
	if n == 0 {
		return true
	}
	last := -1
	for n > 0 {
		c := n & 1
		if last != -1 {
			if c == last {
				return false
			}
		}
		last = c
		n >>= 1
	}
	return true
}
