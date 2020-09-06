package main

import "fmt"

func countOdds(low int, high int) int {
	if (high-low+1)&1 == 0 {
		return (high - low + 1) >> 1
	} else {
		if low&1 == 1 {
			return (high-low+1)>>1 + 1
		} else {
			return (high - low + 1) >> 1
		}
	}
}

func main() {
	fmt.Println(countOdds(0, 1))
}
