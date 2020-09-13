package main

import (
	"fmt"
	"math"
)

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	t := float64(target)
	n := int(math.Ceil((math.Sqrt(8*t+1) - 1) / 2))
	d := n*(n+1)/2 - target
	if d%2 == 0 {
		return n
	} else {
		return n + n%2 + 1
	}
}

func main() {
	target := 2
	fmt.Println(reachNumber(target))
}
