package main

import "fmt"

func minFlips(target string) int {
	res := 0
	last := '0'
	for _, cur := range target {
		if cur == last {
			continue
		} else {
			res++
			last = cur
		}
	}
	return res
}

func main() {
	target := "001011101"
	fmt.Println(minFlips(target))
}
