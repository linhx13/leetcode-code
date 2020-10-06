package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	if len(matrix[0]) == 0 {
		return res
	}
	r_s, c_s := (len(matrix)+1)/2, (len(matrix[0])+1)/2
	m_s := r_s
	if c_s < m_s {
		m_s = c_s
	}
	for s := 0; s < m_s; s++ {
		r, c := len(matrix)-2*s, len(matrix[0])-2*s
		if r == 1 {
			for j := s; j < s+c; j++ {
				res = append(res, matrix[s][j])
			}
			continue
		}
		if c == 1 {
			for i := s; i < s+r; i++ {
				res = append(res, matrix[i][s])
			}
			continue
		}
		for j := s; j < s+c; j++ {
			res = append(res, matrix[s][j])
		}
		for i := s + 1; i < s+r; i++ {
			res = append(res, matrix[i][s+c-1])
		}
		for j := s + c - 2; j >= s; j-- {
			res = append(res, matrix[s+r-1][j])
		}
		for i := s + r - 2; i >= s+1; i-- {
			res = append(res, matrix[i][s])
		}
	}
	return res
}

func main() {
	matrix := [][]int{}

	// matrix = append(matrix, []int{1})
	// matrix = append(matrix, []int{2})
	// matrix = append(matrix, []int{3})
	// matrix = append(matrix, []int{4})

	// matrix = append(matrix, []int{1, 2, 3, 4})

	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})

	// matrix = append(matrix, []int{1, 2, 3, 4})
	// matrix = append(matrix, []int{5, 6, 7, 8})
	// matrix = append(matrix, []int{9, 10, 11, 12})

	// matrix = append(matrix, []int{1, 2, 3, 4})
	// matrix = append(matrix, []int{5, 6, 7, 8})
	// matrix = append(matrix, []int{9, 10, 11, 12})
	// matrix = append(matrix, []int{13, 14, 15, 16})

	fmt.Println(matrix)
	fmt.Println(spiralOrder(matrix))
}
