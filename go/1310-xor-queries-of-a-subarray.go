package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	sums := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			sums[i] = arr[i]
		} else {
			sums[i] = sums[i-1] ^ arr[i]
		}
	}
	res := []int{}
	for i := 0; i < len(queries); i++ {
		x := sums[queries[i][0]] ^ sums[queries[i][1]] ^ arr[queries[i][0]]
		res = append(res, x)
	}
	return res
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 3}, []int{3, 3}}
	fmt.Println(xorQueries(arr, queries))
}
