package main

func setZeroes(matrix [][]int) {
	isCol := false
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0; j < c; j++ {
			matrix[0][j] = 0
		}
	}
	if isCol {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {

}
