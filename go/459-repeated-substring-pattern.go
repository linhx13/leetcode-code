package main

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := n / 2; i >= 1; i-- {
		if n%i == 0 {
			c := n / i
			t := ""
			for j := 0; j < c; j++ {
				t += s[0:i]
			}
			if t == s {
				return true
			}
		}
	}
	return false
}
