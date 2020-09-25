package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	helper(k, n, 1, &[]int{}, &res)
	return res
}

func helper(k int, n int, i int, cur *[]int, res *[][]int) {
	if k == 0 && n == 0 {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}
	if i > 9 {
		return
	}
	*cur = append(*cur, i)
	helper(k-1, n-i, i+1, cur, res)
	*cur = (*cur)[0 : len(*cur)-1]
	helper(k, n, i+1, cur, res)
}

func main() {
	// k, n := 3, 7
	// k, n := 3, 9
	k, n := 4, 1
	fmt.Println(combinationSum3(k, n))
}
