package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	for i := 0; i+10 <= len(s); i++ {
		sub := s[i : i+10]
		m[sub]++
	}
	res := []string{}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
