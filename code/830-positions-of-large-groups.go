package main

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	n := len(S)
	i, j := 0, 0
	for j < n {
		for j < n && S[j] == S[i] {
			j++
		}
		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
		i = j
	}
	return res
}
