package main

import "fmt"

func doSearch(matrix [][]int, low_row int, high_row int, low_col int, high_col int, target int) bool {
	fmt.Println(low_row, high_row, low_col, high_col, target)
	if low_row >= high_row || low_col >= high_col {
		return false
	}
	if target < matrix[low_row][low_col] || target > matrix[high_row-1][high_col-1] {
		return false
	}
	mid_row := low_row + (high_row-low_row)>>1
	mid_col := low_col + (high_col-low_col)>>1
	if target == matrix[mid_row][mid_col] {
		return true
	}
	if target == matrix[low_row][mid_col] {
		return true
	}
	if target < matrix[low_row][mid_col] {
		high_col = mid_col
		return doSearch(matrix, low_row, high_row, low_col, high_col, target)
	}
	if target == matrix[high_row-1][mid_col] {
		return true
	}
	if target > matrix[high_row-1][mid_col] {
		low_col = mid_col + 1
		return doSearch(matrix, low_row, high_row, low_col, high_col, target)
	}
	if target == matrix[mid_row][low_col] {
		return true
	}
	if target < matrix[mid_row][low_col] {
		high_row = mid_row
		return doSearch(matrix, low_row, high_row, low_col, high_col, target)
	}
	if target == matrix[mid_row][high_col-1] {
		return true
	}
	if target > matrix[mid_row][high_col-1] {
		low_row = mid_row + 1
		return doSearch(matrix, low_row, high_row, low_col, high_col, target)
	}
	a := doSearch(matrix, low_row, mid_row, low_col, mid_col, target)
	b := doSearch(matrix, mid_row, high_row, low_col, mid_col, target)
	c := doSearch(matrix, low_row, mid_row, mid_col, high_col, target)
	d := doSearch(matrix, mid_row, high_row, mid_col, high_col, target)
	return a || b || c || d
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return doSearch(matrix, 0, len(matrix), 0, len(matrix[0]), target)
}

func main() {
	matrix := [][]int{}
	matrix = append(matrix, []int{1, 4, 7, 11, 15})
	matrix = append(matrix, []int{2, 5, 8, 12, 19})
	matrix = append(matrix, []int{3, 6, 9, 16, 22})
	matrix = append(matrix, []int{10, 13, 14, 17, 24})
	matrix = append(matrix, []int{18, 21, 23, 26, 30})

	fmt.Println(searchMatrix(matrix, 20))
}
