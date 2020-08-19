package main

import "fmt"

func customSortString(S string, T string) string {
	res := []rune{}
	m := make(map[rune]int)
	for _, c := range T {
		m[c]++
	}
	for _, c := range S {
		for i := 0; i < m[c]; i++ {
			res = append(res, c)
		}
		if _, ok := m[c]; ok {
			delete(m, c)
		}
	}
	for c, v := range m {
		for i := 0; i < v; i++ {
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	S := "cba"
	T := "abcd"
	fmt.Println(customSortString(S, T))
}
