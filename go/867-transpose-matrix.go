package main

func transpose(A [][]int) [][]int {
	res := [][]int{}
	for j := 0; j < len(A[0]); j++ {
		cur := []int{}
		for i := 0; i < len(A); i++ {
			cur = append(cur, A[i][j])
		}
		res = append(res, cur)
	}
	return res
}
