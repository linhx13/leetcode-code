package main

import "fmt"

var m map[int][]string = make(map[int][]string)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	if r, ok := m[n]; ok {
		return r
	}
	res := []string{}
	for i := 0; i <= n-1; i++ {
		if i == 0 {
			for _, s := range generateParenthesis(n - 1) {
				res = append(res, "()"+s)
			}
		} else {
			for _, s := range generateParenthesis(i) {
				if i == n-1 {
					res = append(res, "("+s+")")
				} else {
					for _, t := range generateParenthesis(n - 1 - i) {
						res = append(res, "("+s+")"+t)
					}
				}
			}
		}
	}
	m[n] = res
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
