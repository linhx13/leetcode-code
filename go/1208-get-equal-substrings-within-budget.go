package main

import (
	"fmt"
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	diff := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], t[i])
		diff[i] = int(math.Abs(float64(int(s[i]) - int(t[i]))))
	}
	i := 0
	res := 0
	for j := 0; j < len(s); j++ {
		maxCost -= diff[j]
		for maxCost < 0 {
			maxCost += diff[i]
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}

func main() {
	s, t := "abcd", "bcdf"
	maxCost := 3
	fmt.Println(equalSubstring(s, t, maxCost))
}
