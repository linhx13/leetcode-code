package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := []int{}
	for _, q := range queries {
		A[q[1]] += q[0]
		sum := 0
		for _, n := range A {
			if n%2 == 0 {
				sum += n
			}
		}
		res = append(res, sum)
	}
	return res
}
