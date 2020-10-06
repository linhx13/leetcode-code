package main

import "fmt"

func numSub(s string) int {
	res := 0
	var last byte
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != last {
			if last == '1' {
				if cnt&1 == 1 {
					res += (cnt + 1) >> 1 * cnt
				} else {
					res += cnt >> 1 * (cnt + 1)
				}
			}
			cnt = 0
		}
		last = s[i]
		cnt++
	}
	if last == '1' {
		if cnt&1 == 1 {
			res += (cnt + 1) >> 1 * cnt
		} else {
			res += cnt >> 1 * (cnt + 1)
		}
	}
	return res % (1e9 + 7)
}

func main() {
	s := "101011"
	fmt.Println(numSub(s))
}
