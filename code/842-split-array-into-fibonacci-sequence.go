package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitIntoFibonacci(S string) []int {
	res := []int{}
	helper(S, 0, &[]int{}, &res)
	return res
}

func helper(s string, idx int, buf *[]int, res *[]int) {
	limit := (1<<31 - 1)
	for i := idx + 1; i < len(s); i++ {
		if i-idx > 1 && s[idx] == '0' {
			continue
		}
		a, _ := strconv.Atoi(s[idx:i])
		for j := i + 1; j < len(s); j++ {
			if j-i > 1 && s[i] == '0' {
				continue
			}
			b, _ := strconv.Atoi(s[i:j])
			c := a + b
			if !(a <= limit && b <= limit && c <= limit) {
				continue
			}
			sc := strconv.Itoa(c)
			if strings.HasPrefix(s[j:len(s)], sc) {
				if s[j:len(s)] == sc {
					*res = []int{}
					*res = append(*res, (*buf)...)
					*res = append(*res, a)
					*res = append(*res, b)
					*res = append(*res, c)
					return
				} else {
					*buf = append(*buf, a)
					helper(s, i, buf, res)
					*buf = (*buf)[0 : len(*buf)-1]
				}
			}
		}
	}
}

func main() {
	// S := "11235813"
	// S := "112358130"
	// S := "0123"
	S := "1101111"
	// S := "123456579"
	// S := "0000"
	// S := "1123"
	fmt.Println(splitIntoFibonacci(S))
}
