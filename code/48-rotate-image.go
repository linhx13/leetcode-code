package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
}

func main() {
	matrix := [][]int{}
	// matrix = append(matrix, []int{1, 2, 3})
	// matrix = append(matrix, []int{4, 5, 6})
	// matrix = append(matrix, []int{7, 8, 9})

	// matrix = append(matrix, []int{5, 1, 9, 11})
	// matrix = append(matrix, []int{2, 4, 8, 10})
	// matrix = append(matrix, []int{13, 3, 6, 7})
	// matrix = append(matrix, []int{15, 14, 12, 16})

	// matrix = append(matrix, []int{1})

	matrix = append(matrix, []int{1, 2})
	matrix = append(matrix, []int{3, 4})
	rotate(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
}
