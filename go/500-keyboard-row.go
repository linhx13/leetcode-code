package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	m := map[string]int{
		"q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
		"a": 2, "s": 2, "d": 2, "f": 2, "g": 2, "h": 2, "j": 2, "k": 2, "l": 2,
		"z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3}
	var res []string
	for _, w := range words {
		row := 0
		valid := true
		for _, c := range []rune(w) {
			k := string(c)
			k = strings.ToLower(k)
			if row == 0 {
				row = m[k]
			} else if row != m[k] {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, w)
		}
	}
	return res
}

func main() {
	// words := []string{"Hello", "Alaska", "Dad", "Peace"}
	words := []string{"Aasdfghjkl", "Qwertyuiop", "zZxcvbnm"}
	fmt.Println(findWords(words))
}
