package main

import "fmt"

func reverseStr(s string, k int) string {
	runes := []rune(s)
	for idx := 0; idx < len(s); idx += 2 * k {
		start, end := 0, 0
		if idx+k > len(s) {
			start, end = idx, len(s)-1
		} else if idx+k <= len(s) {
			start, end = idx, idx+k-1
		}
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseStr("abcdefghi", 8))
}
