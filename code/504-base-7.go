package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToBase7(num int) string {
	pos := true
	if num < 0 {
		num = -num
		pos = false
	}
	s := make([]string, 0)
	for {
		a, b := num/7, num%7
		s = append(s, strconv.Itoa(b))
		num = a
		if num == 0 {
			break
		}
	}
	if !pos {
		s = append(s, "-")
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, "")
}

func main() {
	fmt.Println(convertToBase7(-100))
}
