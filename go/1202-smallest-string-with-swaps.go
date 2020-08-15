package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}

	// union
	for _, pair := range pairs {
		root[find(pair[0], root)] = find(pair[1], root)
	}

	ids := make([][]int, n)
	chars := make([][]rune, n)
	for i, c := range []rune(s) {
		p := find(i, root)
		ids[p] = append(ids[p], i)
		chars[p] = append(chars[p], c)
	}

	res := []rune(s)
	for idx := 0; idx < n; idx++ {
		sort.Slice(chars[idx], func(i, j int) bool {
			return chars[idx][i] < chars[idx][j]
		})
		for i := 0; i < len(ids[idx]); i++ {
			res[ids[idx][i]] = chars[idx][i]
		}
	}
	return string(res)
}

func find(x int, root []int) int {
	if root[x] != x {
		root[x] = find(root[x], root)
	}
	return root[x]
}

func main() {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}}
	fmt.Println(smallestStringWithSwaps(s, pairs))
}
