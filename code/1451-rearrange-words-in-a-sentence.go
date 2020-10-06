package main

import (
	"fmt"
	"sort"
	"strings"
)

type pair struct {
	Word string
	Idx  int
}

func arrangeWords(text string) string {
	words := strings.Fields(text)
	pairs := []pair{}
	for i := 0; i < len(words); i++ {
		pairs = append(pairs, pair{Word: strings.ToLower(words[i]), Idx: i})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if len(pairs[i].Word) != len(pairs[j].Word) {
			return len(pairs[i].Word) < len(pairs[j].Word)
		} else {
			return pairs[i].Idx < pairs[j].Idx
		}
	})

	res := []string{}
	for i, p := range pairs {
		if i == 0 {
			res = append(res, strings.Title(p.Word))
		} else {
			res = append(res, p.Word)
		}
	}
	return strings.Join(res, " ")
}

func main() {
	text := "Keep calm and code on"
	fmt.Println(arrangeWords(text))
}
