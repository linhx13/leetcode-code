package main

func findDiagonalOrder(matrix [][]int) []int {
	res := []int{}
	n := len(matrix)
	if n == 0 {
		return res
	}
	m := len(matrix[0])
	dir := 1
	i, j := 0, 0
	for {
		res = append(res, matrix[i][j])
		if i == n-1 && j == m-1 {
			break
		}
		if dir == 1 {
			if j+1 >= m {
				dir = 0
				i = i + 1
			} else if i-1 < 0 {
				dir = 0
				j = j + 1
			} else {
				i--
				j++
			}
		} else {
			if i+1 >= n {
				dir = 1
				j = j + 1
			} else if j-1 < 0 {
				dir = 1
				i = i + 1
			} else {
				i++
				j--
			}
		}
	}
	return res
}
