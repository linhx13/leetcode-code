package main

import "fmt"

func findTheLongestSubstring(s string) int {
	vowels := []byte("aeiou")
	idx := make(map[int]int)
	idx[0] = -1
	state := 0
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < 5; j++ {
			if s[i] == vowels[j] {
				state ^= 1 << j
			}
		}
		if _, ok := idx[state]; !ok {
			idx[state] = i
		}
		if i-idx[state] > res {
			res = i - idx[state]
		}
	}
	return res
}

func main() {
	// s := "eleetminicoworoep"
	s := "leetcodeisgreat"
	fmt.Println(findTheLongestSubstring(s))
}
