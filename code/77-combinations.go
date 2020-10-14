package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	helper(n, k, 1, &[]int{}, &res)
	return res
}

func helper(n int, k int, cur int, buf *[]int, res *[][]int) {
	if k == 0 {
		tmp := make([]int, len(*buf))
		copy(tmp, *buf)
		*res = append(*res, tmp)
		return
	}
	*buf = append(*buf, cur)
	helper(n, k-1, cur+1, buf, res)
	*buf = (*buf)[0 : len(*buf)-1]
	if n-cur >= k {
		helper(n, k, cur+1, buf, res)
	}
}

func main() {
	// n, k := 4, 2
	n, k := 1, 1
	fmt.Println(combine(n, k))
}
