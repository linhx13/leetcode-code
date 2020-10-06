package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	p := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= g[p] {
			res++
			p++
			if p >= len(g) {
				break
			}
		}
	}
	return res
}
