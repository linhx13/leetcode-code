package main

import "fmt"

func longestWord(words []string) string {
	set := make(map[string]bool)
	for _, w := range words {
		set[w] = true
	}
	res := ""
	for _, w := range words {
		isIn := true
		for i := 1; i < len(w); i++ {
			if !set[w[:i]] {
				isIn = false
				break
			}
		}
		if isIn {
			if res == "" || len(w) > len(res) {
				res = w
			} else if len(w) == len(res) && w < res {
				res = w
			}
		}
	}
	return res
}

func main() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(words))
}
