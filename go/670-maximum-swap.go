package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	maxIdx := make([]int, len(s))
	maxIdx[len(s)-1] = len(s) - 1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[maxIdx[i+1]] {
			maxIdx[i] = i
		} else {
			maxIdx[i] = maxIdx[i+1]
		}
	}
	for i := 0; i < len(s); i++ {
		if maxIdx[i] != i && s[i] != s[maxIdx[i]] {
			s[i], s[maxIdx[i]] = s[maxIdx[i]], s[i]
			break
		}
	}
	res, _ := strconv.Atoi(string(s))
	return res
}

func main() {
	num := 2736
	// num := 9973
	// num := 0
	// num := 98368
	// num := 1993
	fmt.Println(maximumSwap(num))
}
