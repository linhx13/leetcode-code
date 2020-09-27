package main

func numSubmat(mat [][]int) int {
	r := len(mat)
	c := len(mat[0])

	for i := 1; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 1 {
				mat[i][j] = mat[i-1][j] + 1
			}
		}
	}

	res := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			min := mat[i][j]
			res += mat[i][j]
			for k := j + 1; k < c; k++ {
				if mat[i][k] == 0 {
					break
				}
				if mat[i][k] < min {
					min = mat[i][k]
				}
				res += min
			}
		}
	}
	return res
}
