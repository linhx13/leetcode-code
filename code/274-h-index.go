package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i := 0; i < len(citations); i++ {
		if i >= citations[i] {
			return i
		}
	}
	return len(citations)
}

func main() {
	// citations := []int{3, 3, 6, 1, 5}
	citations := []int{3, 0, 6, 1, 5}
	// citations := []int{}
	// citations := []int{100}
	fmt.Println(hIndex(citations))
}
