package main

import "fmt"

func numberOfSubstrings(s string) int {
	const INT_MAX = int(^uint(0) >> 1)
	l := make(map[byte]int)
	l['a'] = -1
	l['b'] = -1
	l['c'] = -1
	res := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		l[ch] = i
		min := INT_MAX
		for _, v := range l {
			if v < min {
				min = v
			}
		}
		res += min + 1
	}
	return res
}

func main() {
	s := "abcabc"
	fmt.Println(numberOfSubstrings(s))
}
