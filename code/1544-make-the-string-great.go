package main

import (
	"fmt"
	"math"
)

func makeGood(s string) string {
	res := []rune{}
	diff := int(float64(math.Abs(float64('a' - 'A'))))
	for _, c := range s {
		if len(res) > 0 && int(float64(math.Abs(float64(res[len(res)-1]-c)))) == diff {
			res = res[0 : len(res)-1]
		} else {
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	s := "leEeetcode"
	fmt.Println(makeGood(s))
}
