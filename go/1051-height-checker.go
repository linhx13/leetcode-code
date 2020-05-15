package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	n := 0
	for i := 0; i < len(heights); i++ {
		if sorted[i] != heights[i] {
			n++
		}
	}
	return n
}

func main() {
	heights := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))
}
