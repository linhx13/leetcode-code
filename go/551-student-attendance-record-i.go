package main

func checkRecord(s string) bool {
	a, l := 0, 0
	for _, c := range s {
		if c == 'A' {
			a++
			if a > 1 {
				return false
			}
			l = 0
		} else if c == 'L' {
			l++
			if l > 2 {
				return false
			}
		} else {
			l = 0
		}
	}
	return true
}
