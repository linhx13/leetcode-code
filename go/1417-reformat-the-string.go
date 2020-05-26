package main

import "fmt"

func reformat(s string) string {
	digits, letters := []rune{}, []rune{}
	for _, r := range s {
		if '0' <= r && r <= '9' {
			digits = append(digits, r)
		} else {
			letters = append(letters, r)
		}
	}
	diff := len(digits) - len(letters)
	res := []rune{}
	var x, y []rune
	if diff == 1 {
		res = append(res, digits[0])
		x, y = letters[:], digits[1:]
	} else if diff == -1 {
		res = append(res, letters[0])
		x, y = digits[:], letters[1:]
	} else if diff == 0 {
		x, y = digits[:], letters[:]
	} else {
		return ""
	}
	for i := 0; i < len(x); i++ {
		res = append(res, x[i], y[i])
	}
	return string(res)
}

func main() {
	// s := "a0b1c2"
	// s := "leetcode"
	// s := "1229857369"
	// s := "covid2019"
	s := "a0"
	fmt.Println(reformat(s))
}
