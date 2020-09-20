package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	sort.Strings(words)
	res := 0
	for i := 0; i < len(words)-1; i++ {
		if !strings.HasPrefix(words[i+1], words[i]) {
			res += len(words[i]) + 1
		}
	}
	res += len(words[len(words)-1]) + 1
	return res
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}
