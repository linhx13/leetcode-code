package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	res := []string{}
	for i, x := range words {
		hit := false
		for j, y := range words {
			if i != j && strings.Contains(y, x) {
				hit = true
				break
			}
		}
		if hit {
			res = append(res, x)
		}
	}
	return res
}

func main() {
	// words := []string{"mass", "as", "hero", "superhero"}
	// words := []string{"leetcode", "et", "code"}
	words := []string{"blue", "green", "bu"}
	fmt.Println(stringMatching(words))
}
