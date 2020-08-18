package main

import "fmt"

func numSplits(s string) int {
	c1 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		c1[s[i]]++
	}

	res := 0
	c2 := make(map[byte]int)
	for i := 0; i < len(s)-1; i++ {
		c2[s[i]]++
		c1[s[i]]--
		if c1[s[i]] == 0 {
			delete(c1, s[i])
		}
		if len(c2) == len(c1) {
			res++
		}
	}
	return res
}

func main() {
	// s := "aacaba"
	// s := "abcd"
	// s := "aaaaa"
	s := "acbadbaada"
	fmt.Println(numSplits(s))
}
