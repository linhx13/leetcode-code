package main

import "fmt"

func countBinarySubstrings(s string) int {
	var lens [2]int
	res := 0
	i, l := 0, 0
	for {
		for i < len(s) && s[i] == s[l] {
			i++
		}
		lens[s[l]-'0'] = i - l
		if lens[0] > lens[1] {
			res += lens[1]
		} else {
			res += lens[0]
		}
		if i == len(s) {
			break
		}
		lens[s[i]-'0'] = 0
		l = i
	}
	return res
}

func main() {
	// s := "00110011"
	s := "10101"
	fmt.Println(countBinarySubstrings(s))
}
