package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	res := []string{}
	words := strings.Fields(text)
	for i := 0; i < len(words)-2; i++ {
		if i+1 >= len(words) || i+2 >= len(words) {
			break
		}
		if words[i] == first && words[i+1] == second {
			res = append(res, words[i+2])
		}
	}
	return res
}
