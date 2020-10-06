package main

func maxPower(s string) int {
	var last rune
	max, cur := 0, 0
	for _, c := range s {
		if c != last {
			if last != 0 {
				if cur > max {
					max = cur
				}
			}
			last, cur = 0, 0
		}
		last = c
		cur++
	}
	if last != 0 {
		if cur > max {
			max = cur
		}
	}
	return max
}
