package main

import "fmt"

func maxVowels(s string, k int) int {
	vowels := []byte("aeiou")
	cnts := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		hit := false
		for j := 0; j < len(vowels); j++ {
			if s[i] == vowels[j] {
				hit = true
				break
			}
		}
		if i > 0 {
			cnts[i] = cnts[i-1]
		}
		if hit {
			cnts[i]++
		}
	}
	res := cnts[k-1]
	for i := k; i < len(cnts); i++ {
		if cnts[i]-cnts[i-k] > res {
			res = cnts[i] - cnts[i-k]
		}
	}
	return res
}

func main() {
	// s, k := "abciiidef", 3
	// s, k := "aeiou", 2
	// s, k := "rhythms", 4
	s, k := "tryhard", 4
	fmt.Println(maxVowels(s, k))
}
