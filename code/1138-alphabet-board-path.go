package main

import (
	"fmt"
	"math"
)

func alphabetBoardPath(target string) string {
	pos := make(map[rune][]int)
	for b := 'a'; b <= 'z'; b++ {
		r, c := int((b-'a')/5), int((b-'a')%5)
		pos[b] = []int{r, c}
	}
	res := []rune{}
	r, c := 0, 0
	var last rune
	for _, ch := range target {
		v, _ := pos[ch]
		i, j := v[0], v[1]
		if last == 'z' {
			res = moveRow(r, i, res)
			res = moveCol(c, j, res)
		} else {
			res = moveCol(c, j, res)
			res = moveRow(r, i, res)
		}
		res = append(res, '!')
		r, c = i, j
		last = ch
	}
	return string(res)
}

func moveRow(from int, to int, res []rune) []rune {
	for k := 0; k < int(math.Abs(float64(from-to))); k++ {
		if to > from {
			res = append(res, 'D')
		} else {
			res = append(res, 'U')
		}
	}
	return res
}

func moveCol(from int, to int, res []rune) []rune {
	for k := 0; k < int(math.Abs(float64(from-to))); k++ {
		if to > from {
			res = append(res, 'R')
		} else {
			res = append(res, 'L')
		}
	}
	return res
}

func main() {
	target := "leetzt"
	fmt.Println(alphabetBoardPath(target))
}
