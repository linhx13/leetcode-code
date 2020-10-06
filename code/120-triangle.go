package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	sums := make([]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		if i == 0 {
			sums[i] = triangle[i][i]
		} else {
			for j := i; j >= 0; j-- {
				if j == 0 {
					sums[j] += triangle[i][j]
				} else if j == i {
					sums[j] = triangle[i][j] + sums[j-1]
				} else {
					sums[j] = min(triangle[i][j]+sums[j], triangle[i][j]+sums[j-1])
				}

			}
		}
	}
	res := int(^uint(0) >> 1)
	for i := 0; i < len(triangle); i++ {
		if sums[i] < res {
			res = sums[i]
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	triangle := [][]int{}
	triangle = append(triangle, []int{2})
	triangle = append(triangle, []int{3, 4})
	triangle = append(triangle, []int{6, 5, 7})
	triangle = append(triangle, []int{4, 1, 8, 3})
	fmt.Println(minimumTotal(triangle))
}
