package main

import (
	"fmt"
	"sort"
)

func advantageCount(A []int, B []int) []int {
	sort.Ints(A)
	BB := [][]int{}
	for i, b := range B {
		BB = append(BB, []int{b, i})
	}
	sort.Slice(BB, func(i, j int) bool {
		return BB[i][0] < BB[j][0]
	})
	res := make([]int, len(B))
	for len(A) > 0 {
		a := A[0]
		A = A[1:]
		b := BB[0]
		if a > b[0] {
			BB = BB[1:]
		} else {
			b = BB[len(BB)-1]
			BB = BB[0 : len(BB)-1]
		}
		res[b[1]] = a
	}
	return res
}

func main() {
	// A := []int{2, 7, 11, 15}
	// B := []int{1, 10, 4, 11}
	A := []int{12, 24, 8, 32}
	B := []int{13, 25, 32, 11}
	fmt.Println(advantageCount(A, B))
}
