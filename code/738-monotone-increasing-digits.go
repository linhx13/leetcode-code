package main

import (
	"fmt"
	"strconv"
	"strings"
)

func monotoneIncreasingDigits(N int) int {
	idx := -1
	s := []byte(strconv.Itoa(N))
	sorted := true
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			sorted = false
			idx = i - 1
			break
		}
	}
	if sorted {
		return N
	}
	idx = strings.IndexByte(string(s), s[idx])
	s[idx] = s[idx] - byte(1)
	for i := idx + 1; i < len(s); i++ {
		s[i] = '9'
	}
	i := 0
	for ; i < len(s); i++ {
		if s[i] != '0' {
			break
		}
	}
	s = s[i:len(s)]
	res, _ := strconv.Atoi(string(s))
	return res
}

func main() {
	N := 332
	fmt.Println(monotoneIncreasingDigits(N))
}
