package main

func maxScore(s string) int {
	ones := 0
	for _, c := range s {
		if c == '1' {
			ones++
		}
	}
	max := 0
	left, right := 0, ones
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		} else {
			right--
		}
		if left+right > max {
			max = left + right
		}
	}
	return max
}
