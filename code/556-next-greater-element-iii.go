package main

import (
	"fmt"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	INT_MAX := 0x7fffffff
	s := []byte(strconv.Itoa(n))
	res := -1
	for i := len(s) - 2; i >= 0; i-- {
		x, j := min(s, i)
		if x == 0 {
			continue
		}
		s[i], s[j] = s[j], s[i]
		sort.Slice(s[i+1:len(s)], func(x, y int) bool {
			return s[i+1+x] < s[i+1+y]
		})
		m, _ := strconv.Atoi(string(s))
		if m <= INT_MAX {
			res = m
		}
		break
	}
	return res
}

func min(s []byte, i int) (byte, int) {
	var res byte
	var k int
	for j := i + 1; j < len(s); j++ {
		if s[j] > s[i] && (res == 0 || s[i] < res) {
			res = s[j]
			k = j
		}
	}
	return res, k
}

func main() {
	// n := 12
	// n := 21
	// n := 1234
	// n := 1222333
	n := 1999999999
	// n := 2147483647
	// n := 13
	fmt.Println(nextGreaterElement(n))
}
