package main

import "math"

func imageSmoother(M [][]int) [][]int {
	rows := len(M)
	cols := len(M[0])
	res := [][]int{}
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		res = append(res, row)
	}
	ranges := []int{-1, 0, 1}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cur := 0
			cnt := 0
			for _, a := range ranges {
				for _, b := range ranges {
					row := i + a
					col := j + b
					if row >= 0 && row < rows && col >= 0 && col < cols {
						cur += M[row][col]
						cnt++
					}
				}
			}
			res[i][j] = int(math.Floor(float64(cur) / float64(cnt)))
		}
	}
	return res
}
