package main

import "fmt"

func minSteps(s string, t string) int {
	chars := make(map[rune]bool)
	cs := make(map[rune]int)
	for _, c := range s {
		chars[c] = true
		cs[c] += 1
	}
	ct := make(map[rune]int)
	for _, c := range t {
		chars[c] = true
		ct[c] += 1
	}
	res := 0
	for c, _ := range chars {
		vt, vs := ct[c], cs[c]
		if vt < vs {
			res += vs - vt
		}
	}
	return res
}

func main() {
	s := "leetcode"
	t := "practice"
	fmt.Println(minSteps(s, t))
}
