package main

import "fmt"

func balancedString(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	avg := len(s) / 4
	for _, ch := range []byte{'Q', 'W', 'E', 'R'} {
		if m[ch] <= avg {
			delete(m, ch)
		} else {
			m[ch] = m[ch] - avg
		}
	}
	if len(m) == 0 {
		return 0
	}
	res := int(^uint(0) >> 1)
	for i, l, sz := 0, 0, len(m); i < len(s); i++ {
		ch := s[i]
		if _, ok := m[ch]; !ok {
			continue
		}
		m[ch]--
		if m[ch] == 0 {
			sz--
		}
		for sz == 0 {
			x := s[l]
			l++
			if _, ok := m[x]; !ok {
				continue
			}
			m[x]++
			if m[x] == 1 {
				if i-l+2 < res {
					res = i - l + 2
				}
				sz++
			}
		}
	}
	return res
}

func main() {
	s := "QQQQ"
	fmt.Println(balancedString(s))
}
