package main

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	m := len(rectangle)
	n := len(rectangle[0])
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = rectangle[i][j]
		}
		matrix[i] = row
	}
	return SubrectangleQueries{rectangle: matrix}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}
