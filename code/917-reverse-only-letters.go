package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(S string) string {
	runes := []rune(S)
	i, j := 0, len(runes)-1
	for i < j {
		if !unicode.IsLetter(runes[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(runes[j]) {
			j--
			continue
		}
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func main() {
	// s := "ab-cd"
	// s := "a-bC-dEf-ghIj"
	// s := "Test1ng-Leet=code-Q!"
	s := ""
	fmt.Println(reverseOnlyLetters(s))
}
