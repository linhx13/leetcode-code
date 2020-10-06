package main

import "fmt"

func processQueries(queries []int, m int) []int {
	p := []int{}
	for i := 1; i <= m; i++ {
		p = append(p, i)
	}
	res := []int{}
	for i := 0; i < len(queries); i++ {
		j := 0
		for ; j < len(p); j++ {
			if p[j] == queries[i] {
				break
			}
		}
		res = append(res, j)
		for ; j > 0; j-- {
			p[j] = p[j-1]
		}
		p[0] = queries[i]
	}
	return res
}

func main() {
	// queries := []int{3, 1, 2, 1}
	// m := 5
	// queries := []int{4, 1, 2, 2}
	// m := 4
	queries := []int{7, 5, 5, 8, 3}
	m := 8
	fmt.Println(processQueries(queries, m))
}
