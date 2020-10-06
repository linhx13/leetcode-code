package main

import "fmt"

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	if k == 0 {
		return false
	}
	cnts := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if t[i] == s[i] {
			continue
		}
		diff := int(int(t[i])-int(s[i])) + 26
		cnts[diff%26]++
	}
	for d, c := range cnts {
		if k < 26*(c-1)+d {
			return false
		}
	}
	return true
}

func main() {
	s, t := "input", "ouput"
	k := 9
	// s, t := "abc", "bcd"
	// k := 10
	// s, t := "aab", "bbb"
	// k := 27
	// s, t := "z", "a"
	// k := 1
	fmt.Println(canConvertString(s, t, k))
}
