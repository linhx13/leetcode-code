package main

func longestPalindrome(s string) int {
	freqs := make(map[rune]int)
	for _, c := range s {
		freqs[c] += 1
	}
	res := 0
	odd := 0
	for _, f := range freqs {
		if f&1 == 1 {
			res += f - 1
			odd |= 1
		} else {
			res += f
		}
	}
	res += odd
	return res
}
