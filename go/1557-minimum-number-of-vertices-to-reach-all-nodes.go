package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	cnts := make([]int, n)
	for _, e := range edges {
		cnts[e[1]]++
	}
	res := []int{}
	for i, c := range cnts {
		if c == 0 {
			res = append(res, i)
		}
	}
	return res
}
